.PHONY: critic security vulncheck lint test all

mod:
	go list -m --versions

test:
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

critic:
	gocritic check -enableAll -disable=unnamedResult,unlabelStmt,hugeParam,singleCaseSwitch,builtinShadow,typeAssertChain ./...

security:
	gosec -exclude-dir=mysql,psql -exclude=G103,G115,G401,G501 ./...

vulncheck:
	govulncheck ./...

lint:
	golangci-lint run ./...

all: critic security vulncheck lint test