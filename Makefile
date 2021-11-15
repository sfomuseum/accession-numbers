OS := $(shell uname | tr '[:upper:]' '[:lower:]')

cli:
	env GOOS=darwin GOARCH=amd64 go build -o bin/darwin/test-runner cmd/test-runner/main.go
	env GOOS=linux GOARCH=amd64 go build -o bin/linux/test-runner cmd/test-runner/main.go
	env GOOS=windows GOARCH=amd64 go build -o bin/windows/test-runner cmd/test-runner/main.go

tests:
	bin/$(OS)/test-runner data/*.json

data-docs:
	go run cmd/data-docs/main.go data/*.json > data/README.md
