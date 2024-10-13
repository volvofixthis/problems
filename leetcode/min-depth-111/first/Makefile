.PHONY: escape test

test:
	go test ./...

bench:
	go test -bench=. -benchmem

escape:
	go build -gcflags='-m=3' . |& grep escape
