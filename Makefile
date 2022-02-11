xlsxlang: grammer.peg.go *.go */*.go
	go fmt ./...
	go build -tags debug
	go test -cover

grammer.peg.go: grammer.peg
	peg grammer.peg

.PHONY: setup
setup:
	go install github.com/pointlander/peg
