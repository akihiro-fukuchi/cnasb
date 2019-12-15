GOFILES = `find . -type f -name *.go`

fmt:
	goimports -d -w $(GOFILES)

test: fmt
	go test ./...
	go mod tidy

install:
	go install ./cmd/grev

