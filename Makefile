PACKAGES=$(shell go list ./...)

default: check_deps check_lint build

check_deps:
	@echo "--> Checking deps"
	go mod download

check_lint:
	@echo "--> Running golangci"
	golangci-lint run --verbose

build:
	go build -o ./bin/cosmos-gen *.go

buidl: build

build-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/cosmos-gen *.go

install_tools_macos:
	brew install golangci/tap/golangci-lint
