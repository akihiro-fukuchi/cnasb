GOFILES = `find . -type f -name *.go`

goimports:
	go get golang.org/x/tools/cmd/goimports

fmt: goimports
	goimports -d -w $(GOFILES)

test: fmt
	go test ./...
	go mod tidy

install:
	go install ./cmd/grev

