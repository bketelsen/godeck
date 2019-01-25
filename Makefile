GO_BIN ?= go

install:
	@$(GO_BIN) install -v ./.
	@make tidy

tidy:
ifeq ($(GO111MODULE),on)
	@$(GO_BIN) mod tidy
endif

deps:
	@$(GO_BIN) get github.com/gobuffalo/release
	@$(GO_BIN) get github.com/gobuffalo/shoulders
	@$(GO_BIN) get -t ./...
	@make tidy

build:
	@$(GO_BIN) build -v .
	@make tidy

test:
	@packr2
	@$(GO_BIN) test ./...
	@make tidy

shoulders:
	@shoulders -n godeck -w

ci-deps:
	$(GO_BIN) get -tags -t ./...

ci-test:
	$(GO_BIN) test -tags -race ./...

lint:
	@gometalinter --vendor ./... --deadline=1m --skip=internal
	@make tidy

update:
	@$(GO_BIN) get -u 
	@make tidy
	@make test
	@make install
	@make tidy

release-test:
	@$(GO_BIN) test -tags ${TAGS} -race ./...
	@make tidy

release:
	@make tidy
	@make shoulders
	@release -y -f ./version/version.go
	@make tidy
