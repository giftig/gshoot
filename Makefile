build:
	@go build

test:
	@mkdir -p cov
	@go clean -testcache
	go test -coverprofile=cov/coverage.txt ./...
	@go tool cover -html=cov/coverage.txt -o cov/coverage.html

run:
	@go run main.go

install:
	sudo cp gshoot /usr/local/bin/gshoot
