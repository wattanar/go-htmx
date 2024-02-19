run:
	go run cmd/server.go

build:
	go build -o bin/app cmd/server.go

test: 
	go test ./...