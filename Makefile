run:
	@go run

test: bootstrap
	@mkdir -p cov
	@go clean -testcache
	go test -coverprofile=cov/coverage.txt ./...
	@go tool cover -html=cov/coverage.txt -o cov/coverage.html
