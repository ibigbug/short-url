run: 
	sh ./scripts/start.sh

dev:
	go run main.go middlewares.go handlers.go services.go types.go

test:
	go test -v -race -bench .

ci:
	go test -v ./... -covermode=count -coverprofile=coverage.out
	$(HOME)/gopath/bin/goveralls -coverprofile=coverage.out -service=travis-ci -repotoken $(COVERALLS_TOKEN)
