GO_FILES=$(shell find . -name "*.go" -type f)
BUNDLE=bin
DOCKER_RUN=docker run -it --rm -v "$$PWD":/go/src/github.com/grid-x/commandwrapper -w /go/src/github.com/grid-x/commandwrapper
GOLINT=
GO_VERSION=1.10.2
GOARCH=amd64
GOOS ?= $(shell go env GOOS)

lint:
	${DOCKER_RUN} gridx/golang-tools:${GO_VERSION} bash -c "export GOPATH=/go && golint ${GO_FILES}"

bin/commandwrapper: 
	${DOCKER_RUN} gridx/golang-tools:${GO_VERSION} bash -c "export GOPATH=/go && CGO_ENABLED=0 GOOS=${GOOS} GOARCH=$(GOARCH) go build -o ${BUNDLE}/commandwrapper"