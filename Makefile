run:
	go run ./cmd/app/main.go

build:
	go build -o bin/main ./cmd/app/main.go

test:
	go test -v -cover ./...

.PHONY: run build test
