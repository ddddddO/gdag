lint:
	golangci-lint run

test:
	go clean -testcache
	go test . -race -v -count=1
