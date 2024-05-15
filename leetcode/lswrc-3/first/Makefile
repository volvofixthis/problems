.PHONY: escape test

test:
	go test ./...

escape:
	go build -gcflags='-m=3' . |& grep escape
