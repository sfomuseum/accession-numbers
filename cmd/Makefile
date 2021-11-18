cli:
	env GOOS=darwin GOARCH=amd64 go build -mod vendor -o ../bin/darwin/test-runner test-runner/main.go
	env GOOS=linux GOARCH=amd64 go build -mod vendor -o ../bin/linux/test-runner test-runner/main.go
	env GOOS=windows GOARCH=amd64 go build -mod vendor -o ../bin/windows/test-runner test-runner/main.go

docs:
	go run -mod vendor data-docs/main.go ../data/*.json > ../data/README.md

