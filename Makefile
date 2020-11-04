build:
	@go build -o releases/redirect

dockerbuild:
	docker build -t c4po/redirect .

dockerpush:
	docker push c4po/redirect

fmt:
	go mod tidy
	gofmt -s -w .
# 	golangci-lint run ./

tools:
	GO111MODULE=on go install github.com/golangci/golangci-lint/cmd/golangci-lint
