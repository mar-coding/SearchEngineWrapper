check:
	gofumpt -l -w .
	golangci-lint run --timeout=20m0s

install-lint-tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install mvdan.cc/gofumpt@latest

build-search:
	go build -mod=vendor -ldflags "-X 'github.com/mar-coding/SearchEngineWrapper/main.CommitId=$(shell git rev-parse --short HEAD 2>/dev/null)' -X 'github.com/mar-coding/SearchEngineWrapper/main.Date=$(shell date)'" -o main cmd/main.go

	#for running without vendor
	#go build -mod=mod -ldflags "-X 'github.com/mar-coding/SearchEngineWrapper/main.CommitId=$(shell git rev-parse --short HEAD 2>/dev/null)' -X 'github.com/mar-coding/SearchEngineWrapper/main.Date=$(shell date)'" -o main cmd/main.go
