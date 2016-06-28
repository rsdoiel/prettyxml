#
# Simple Makefile
#

build:
	go build -o bin/prettyxml cmds/prettyxml/prettyxml.go 

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi

install:
	env GOBIN=$(HOME)/bin go install cmds/prettyxml/prettyxml.go

release:
	./mk-release.sh
