.PHONY: test
run/tests:
	go test -v ./...
	go test -v -race $(shell go list ./... | grep -v vendor)
