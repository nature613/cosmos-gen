PACKAGES=$(shell go list ./...)

default: check_deps check_lint build

check_deps:
	@echo "--> Checking deps"
	go mod download

check_lint:
	@echo "--> Running golangci"
	golangci-lint run --verbose

build:
	packr build -o ./bin/cosmos-gen *.go

buidl: build

build-linux:
	go get -u github.com/gobuffalo/packr/packr
	mkdir -p bin
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 packr build && mv cosmos-gen ./bin/cosmos-gen

install_tools_macos:
	brew install golangci/tap/golangci-lint
	go get -u github.com/gobuffalo/packr/packr
