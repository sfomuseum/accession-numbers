GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")

cli:
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/twilio-handler cmd/twilio-handler/main.go
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/flatten-definition cmd/flatten-definition/main.go

lambda:
	@make lambda-twilio-handler

lambda-twilio-handler:
	if test -f bootstrap; then rm -f bootstrap; fi
	if test -f twilio-handler.zip; then rm -f twilio-handler.zip; fi
	GOARCH=arm64 GOOS=linux go build -mod $(GOMOD) -ldflags="-s -w" -tags lambda.norpc -o bootstrap cmd/twilio-handler/main.go
	zip twilio-handler.zip bootstrap
	rm -f bootstrap
