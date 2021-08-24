.DEFAULT_GOAL := all

GO=$(shell which go)
DISTVER=$(shell git describe --always --dirty --long --tags)
PKG=$(shell head -1 go.mod | sed 's/^module //')

all: dist

test:
	$(GO) test -v ./...
#	$(GO) test
#	$(GO) test -v ./dtc/...

dist:
	$(GO) build -ldflags "-X main.Version=$(DISTVER) -X $(PKG)/web.Version=$(DISTVER) -X $(PKG)/dtc.Version=$(DISTVER)"

install:
	$(GO) install -ldflags "-X main.Version=$(DISTVER) -X $(PKG)/web.Version=$(DISTVER) -X $(PKG)/dtc.Version=$(DISTVER)" .

dev:
	$(GO) run -ldflags "-X main.Version=$(DISTVER) -X $(PKG)/web.Version=$(DISTVER) -X $(PKG)/dtc.Version=$(DISTVER)" . 2>>sc-dtc-client.err

race:
	$(GO) run -ldflags "-X main.Version=$(DISTVER) -X $(PKG)/web.Version=$(DISTVER) -X $(PKG)/dtc.Version=$(DISTVER)" --race . 2>>sc-dtc-client.err

upgrade:
	$(GO) get -u && $(GO) mod tidy

goformat:
	gofmt -s -w .

