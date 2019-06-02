GO_BIN ?= go
OS := $(shell uname -s)

install: deps
	make build
	$(GO_BIN) install -v .

tidy:
	$(GO_BIN) mod tidy

deps:
	$(GO_BIN) get -u github.com/gobuffalo/packr/v2/packr2
	packr2 clean
	make tidy

build:
ifeq ($(OS), Darwin)
	packr2 build -v .
endif
ifeq ($(OS), Linux)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 packr2 build -v .
endif

buidl: build

# --- Linting
lint:
	golangci-lint run --verbose

install-lint:
ifeq ($(OS), Darwin)
	brew install golangci/tap/golangci-lint
endif
ifeq ($(OS), Linux)
	GO111MODULE=on $(GO_BIN) get github.com/golangci/golangci-lint/cmd/golangci-lint@692dacb773b703162c091c2d8c59f9cd2d6801db
endif
# ---
