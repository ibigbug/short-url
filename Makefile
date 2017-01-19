run: build
	./bin/short-url

build:
	sh ./scripts/build.sh

dev:
	go run main.go middlewares.go handlers.go services.go types.go

test:
	go test -v -race -bench .
