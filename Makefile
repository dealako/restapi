SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

APP_NAME=restapi
BINARY=$(APP_NAME)

VERSION=1.0.0
BUILD_TIME=`date +%FT%T%z`
VERSION := $(shell sh -c 'git describe --always --tags')
BRANCH := $(shell sh -c 'git rev-parse --abbrev-ref HEAD')
COMMIT := $(shell sh -c 'git rev-parse --short HEAD')

LDFLAGS=-ldflags "-s -X main.name=$(APP_NAME) -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.branch=$(BRANCH) -X main.buildDate=$(BUILD_TIME)"

GO_META_LINTER=../../../../bin/gometalinter

.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(SOURCES)
	go build ${LDFLAGS} -o ${BINARY} main.go

.PHONY: install
install:
	go install ${LDFLAGS} ./...

.PHONY: clean
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: install-deps
install-deps: $(GO_META_LINTER)
	dep ensure

$(GO_META_LINTER):
	@echo "Downloading gometalinter..."
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

.PHONY: analyze
analyze: install-deps
	@echo "Analyzing code..."
	-gometalinter --disable=gotype --enable=gofmt --enable=goimports --enable=unused --deadline=2m --vendor ./...

imports:
	goimports -w *.go

fmt:
	gofmt -w -s *.go
