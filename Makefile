GO_FILES=$(shell find . -name "*.go" -type f)
BUNDLE=bin
DOCKER_RUN=docker run -it --rm -v "$$PWD":/go/src/github.com/grid-x/commandwrapper -w /go/src/github.com/grid-x/commandwrapper
GO_VERSION=1.12
GO_TOOLS=gridx/golang-dev:${GO_VERSION}.latest-linux-amd64
GOARCH=amd64
GOOS ?= $(shell go env GOOS)
CI_DEPS=apt-get update && apt-get install -y procps

lint:
	${DOCKER_RUN} ${GO_TOOLS} bash -c "export GOPATH=/go && golint ${GO_FILES}"

test:
	${DOCKER_RUN} ${GO_TOOLS} bash -c "${CI_DEPS} && export GOPATH=/go && go test"

# Build bin to debug on amd64
bin/commandwrapper: 
	${DOCKER_RUN} ${GO_TOOLS} bash -c "export GOPATH=/go && CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags=\"-w -s\" -o ${BUNDLE}/commandwrapper"
