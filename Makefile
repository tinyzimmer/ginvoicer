NAME    = ginvoicer
BINDIR ?= $(CURDIR)/bin
GOBIN  ?= $(shell go env GOPATH)/bin

build:
	cd cmd/$(NAME) && go build -o $(BINDIR)/$(NAME) .

install: build
	cp $(BINDIR)/$(NAME) $(GOBIN)/$(NAME)

GOLANGCI_LINT    ?= $(BINDIR)/golangci-lint
GOLANGCI_VERSION ?= v1.41.1
$(GOLANGCI_LINT):
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$(PWD)/bin" $(GOLANGCI_VERSION)

lint: $(GOLANGCI_LINT)
	"$(GOLANGCI_LINT)" run -v

GOX ?= $(GOBIN)/gox
$(GOX):
	GO111MODULE=off go get github.com/mitchellh/gox

DIST ?= $(CURDIR)/dist
COMPILE_TARGETS ?= "darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 windows/amd64"
COMPILE_OUTPUT  ?= "$(DIST)/{{.Dir}}_{{.OS}}_{{.Arch}}"
LDFLAGS ?= -s -w
dist-ginvoicer: $(GOX)
	mkdir -p "$(DIST)"
	cd cmd/ginvoicer && \
		CGO_ENABLED=0 $(GOX) -osarch=$(COMPILE_TARGETS) -output=$(COMPILE_OUTPUT) -ldflags="$(LDFLAGS)"
	upx -9 $(DIST)/*
	cd dist && sha256sum * > ginvoicer.sha256sum