xlsxlang: grammer.peg.go *.go */*.go
	go fmt .
	go build
	go test -cover

grammer.peg.go: grammer.peg
	peg grammer.peg

.PHONY: setup
setup:
	go install github.com/pointlander/peg
