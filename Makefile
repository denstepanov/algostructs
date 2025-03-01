.DEFAULT_GOAL := build

build:
	go fmt ./... && go vet ./... && go build src/algostructs.go
run:
	go run src/algostructs.go
test:
	go test -v ./...