#
# Simple Makefile
#
build:
	goimport -w cmds/prettyxml/prettyxml.go
	gofmt -w cmds/prettyxml/prettyxml.go
	go build
	go build -o bin/prettyxml cmds/prettyxml/prettyxml.go

test:
	go test

save:
	./mk-website.bash
	git commit -am "quick save"
	git push origin master

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -f prettyxml-binary-release.zip ]; then rm -f prettyxml-binary-release.zip; fi

install:
	env GOBIN=$(HOME)/bin go install cmds/prettyxml/prettyxml.go

release:
	./mk-release.bash

publish:
	./mk-website.bash
	./publish.bash
