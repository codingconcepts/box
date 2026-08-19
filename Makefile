build:
	go build .
	- mv box ~/dev/bin

test:
	go test ./... -v -cover