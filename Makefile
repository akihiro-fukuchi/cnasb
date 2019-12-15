GOFILES = `find . -type f -name *.go`

goimports:
	GO111MODULE=off GOBIN=$(PWD)/bin go get golang.org/x/tools/cmd/goimports

fmt: goimports
	bin/goimports -d -w $(GOFILES)

test: fmt
	go test ./...
	go mod tidy

install:
	go install ./cmd/grev

