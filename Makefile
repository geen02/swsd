# borrow from https://github.com/genuinetools/udict/blob/master/basic.mk
NAME := swsd
PKG := github.com/geen02/$(NAME)
GO_SOURCE := $(shell pwd)/cmd/$(NAME)
# Go parameters
GO := go
GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GO_DEP := dep ensure

GOOS := windows
# GOOS := darwin
GOARCH := amd64

BINARY_NAME=$(NAME)
BINARY_UNIX=$(BINARY_NAME)_unix
CGO_ENABLED := 1

BUILDTAGS := seccomp


# Set an output prefix, which is the local directory if not specified
PREFIX?=$(shell pwd)

# Set the build dir, where built cross-compiled binaries will be output
BUILDDIR := ${PREFIX}/cross

# Populate version variables
# Add to compile time flags
VERSION := $(shell cat VERSION.txt)
GITCOMMIT := $(shell git rev-parse --short HEAD)
GITUNTRACKEDCHANGES := $(shell git status --porcelain --untracked-files=no)
ifneq ($(GITUNTRACKEDCHANGES),)
	GITCOMMIT := $(GITCOMMIT)-dirty
endif
ifeq ($(GITCOMMIT),)
		GITCOMMIT := ${GITHUB_SHA}
endif
CTIMEVAR=-X $(PKG)/version.GITCOMMIT=$(GITCOMMIT) -X $(PKG)/version.VERSION=$(VERSION)
GO_LDFLAGS=-ldflags "-w $(CTIMEVAR)"
GO_LDFLAGS_STATIC=-ldflags "-w $(CTIMEVAR) -extldflags -static"


all: test build
build-win:
	mkdir -p $(BUILDDIR)
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) $(GOBUILD) \
		 -o $(BUILDDIR)/$(NAME)-$(GOOS)-$(GOARCH).exe \
		 -a -tags "$(BUILDTAGS) static_build netgo" \
		 -installsuffix netgo ${GO_LDFLAGS_STATIC} $(GO_SOURCE)
	md5 $(BUILDDIR)/$(NAME)-$(GOOS)-$(GOARCH).exe > $(BUILDDIR)/$(NAME)-$(GOOS)-$(GOARCH).md5
	shasum $(BUILDDIR)/$(NAME)-$(GOOS)-$(GOARCH).exe > $(BUILDDIR)/$(NAME)-$(GOOS)-$(GOARCH).sha256

build:
	mkdir -p $(BUILDDIR)
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) $(GOBUILD) \
		 -o $(BUILDDIR)/$(NAME)-$(GOOS)-$(GOARCH) \
		 -a -tags "$(BUILDTAGS) static_build netgo" \
		 -installsuffix netgo ${GO_LDFLAGS_STATIC} $(GO_SOURCE)
	md5 $(BUILDDIR)/$(NAME)-$(GOOS)-$(GOARCH) > $(BUILDDIR)/$(NAME)-$(GOOS)-$(GOARCH).md5
	shasum $(BUILDDIR)/$(NAME)-$(GOOS)-$(GOARCH) > $(BUILDDIR)/$(NAME)-$(GOOS)-$(GOARCH).sha256

test:
	$(GOTEST) -v $(shell pwd)/pkg/swsdLib/*.go

clean:
	$(GOCLEAN)
	rm -rf $(BUILDDIR)

run:

deps:
	$(GO_DEP)

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
docker-build:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v

