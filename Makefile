GO_BIN ?= go

install: deps
	make build
	$(GO_BIN) install -v .

tidy:
	$(GO_BIN) mod tidy

deps:
	$(GO_BIN) get github.com/gobuffalo/packr/v2/packr2
	packr2 clean
	make tidy

lint:
	golangci-lint run --verbose

build:
	packr2 build -v .

buidl: build

build-linux:
	$(GO_BIN) get -u github.com/gobuffalo/packr/v2/packr2
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 packr2 build

install-tools-macos:
	brew install golangci/tap/golangci-lint

install-tools-linux:
	GO111MODULE=on $(GO_BIN) get github.com/golangci/golangci-lint/cmd/golangci-lint@692dacb773b703162c091c2d8c59f9cd2d6801db
