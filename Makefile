OS := $(shell uname | tr '[:upper:]' '[:lower:]')

tests:
	bin/$(OS)/test-runner data/*.json

docs:
	cd cmd && make docs && cd -

index:
	cd data && ls -a *.json > index.txt && cd -

release:
	@make docs
	@make index
