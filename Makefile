NAME    = ginvoicer
BINDIR ?= $(CURDIR)/bin

build:
	cd cmd/$(NAME) && go build -o $(BINDIR)/$(NAME) .
