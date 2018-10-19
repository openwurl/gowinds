# GoWinds Makefile
#///////////////////////#
#///      DEFS       ///#
#///////////////////////#
# Don't ask, for to understand it is to look 
# into the void and know the void is not only 
#looking back but also reading your emails.
SHELL=/bin/bash -e -o pipefail

# ENV Vars defaults
GOOS ?= darwin
GOARCH ?= amd64
VERSION ?= v0.1

#///////////////////////#
#///      OUTPUT     ///#
#///////////////////////#

RED=\033[0;31m
GREEN=\033[0;32m
YELLOW=\033[01;33m
BLUE=\033[0;34m
LBLUE=\033[01;34m
ORANGE=\033[0;33m
PURPLE=\033[0;35m
LCYAN=\033[1;36m
NC=\033[0m

#///////////////////////#
#///      TARGETS    ///#
#///////////////////////#

.PHONY: help

help:           ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY: test cover build buildall clean

travis: clean test build

test:	## Run test with coverage
	go test -v -race -cover -coverprofile=cov.out

cover:	## Open coverage
	go tool cover --html=cov.out

build:	## Build for testing
	go build -o target/gowinds

clean:	## Clean build
	@go clean
	@rm -rf target
	@rm -rf cov.out