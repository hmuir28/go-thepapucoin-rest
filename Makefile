build:
	@go build -o ./bin/go-thepapucoin-rest

run: build
	@./bin/go-thepapucoin-rest

test:
	go test -v ./...
