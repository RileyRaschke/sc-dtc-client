.DEFAULT_GOAL := dist

MAIN=sc-dtc-client

GO=$(shell which go)
DISTVER=$(shell git describe --always --dirty --long --tags)
PKG=github.com/RileyR387/sc-dtc-client

dist:
	$(GO) build -ldflags "-X main.Version=$(DISTVER) -X $(PKG)/web.Version=$(DISTVER) -X $(PKG)/dtc.Version=$(DISTVER)"

dev:
	$(GO) run -ldflags "-X main.Version=$(DISTVER) -X $(PKG)/web.Version=$(DISTVER) -X $(PKG)/dtc.Version=$(DISTVER)" . 2>>sc-dtc-client.err

test:
	$(GO) test

race:
	$(GO) run -ldflags "-X main.Version=$(DISTVER) -X $(PKG)/web.Version=$(DISTVER) -X $(PKG)/dtc.Version=$(DISTVER)" --race . 2>>sc-dtc-client.err

goformat:
	gofmt -s -w .

