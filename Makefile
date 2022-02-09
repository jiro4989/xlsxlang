xlsxlang:
	go build

grammer.peg.go: grammer.peg
	peg grammer.peg

.PHONY: setup
setup:
	go install github.com/pointlander/peg@latest
