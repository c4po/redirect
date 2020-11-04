package main

import (
    "fmt"
    "net/http"
    "regexp"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	reg string
	targetPattern string
)

func main() {
	kingpin.Flag("regex", "").Envar("URL_REG").Default("/(.+)/(.+)").StringVar(&reg)
	kingpin.Flag("pattern", "").Envar("URL_PATTERN").Default("").StringVar(&targetPattern)
	kingpin.Parse()

	fmt.Println (reg, targetPattern)
    http.HandleFunc("/", RedirectServer)
    http.ListenAndServe(":8080", nil)
}

func RedirectServer(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile(reg)
	match := re.FindStringSubmatch(r.URL.Path)
	fmt.Println (match)
	targetUrl := fmt.Sprintf(targetPattern,match[1],match[2])
    http.Redirect(w, r, targetUrl, 301)
}
