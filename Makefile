NAME    = ginvoicer
BINDIR ?= $(CURDIR)/bin
GOBIN  ?= $(shell go env GOPATH)/bin

build:
	cd cmd/$(NAME) && go build -o $(BINDIR)/$(NAME) .

install: build
	cp $(BINDIR)/$(NAME) $(GOBIN)/$(NAME)