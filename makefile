.PHONY: defalt run build test docs clean
#variables
APP_NAME=go-portunities

#tasks
default: run-with-docs

run:
			@go run main.go
run-with-docs:
			@swag init
			@go run main.go
install:
			@go mod tidy
build: 
			@go build -o $(APP_NAME) main.go
test:
			@go test ./ ...
docs:
			@swag init
clean:
			@rm -f $(APP_NAME)
			@rm -rf ./docs