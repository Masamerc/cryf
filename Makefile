APP_NAME:=cryf
VERSION:=0.0.1

dev_build:
	@echo "build dev binary under target/dev"
	@go build -o target/dev/$(APP_NAME) ./dev/main.go

test:
	@echo "run test"
	@go test -v ./...

help:
	@echo "Available make commands:"
	@echo "- make dev_build:    build dev binary under target/dev"
	@echo "- make test:         run unit test against all source code"
	@echo "- make help:         show this help message"