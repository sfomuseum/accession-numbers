OS := $(shell uname | tr '[:upper:]' '[:lower:]')

tests:
	bin/$(OS)/test-runner data/*.json

data-docs:
	go run cmd/data-docs/main.go data/*.json > data/README.md
