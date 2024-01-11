BINARY_NAME=go-networking

build:
	@go build -o bin/$(BINARY_NAME) -v cmd/${BINARY_NAME}/main.go

run: build
	@./bin/$(BINARY_NAME)

clean:
	@rm -rf ./bin