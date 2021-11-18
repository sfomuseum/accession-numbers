OS := $(shell uname | tr '[:upper:]' '[:lower:]')

tests:
	bin/$(OS)/test-runner data/*.json
