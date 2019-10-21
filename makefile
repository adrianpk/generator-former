# Vars
PROD_TAG=v0.0.1

# Misc
BINARY_NAME=mw
BINARY_UNIX=$(BINARY_NAME)_unix

# Env
INSTALL_PATH=$(HOME)/Dev/go/general/bin

all: test build

build:
	go build -o ./bin/$(BINARY_NAME) main.go metadata.go migration.go

build-linux:
	CGOENABLED=0 GOOS=linux GOARCH=amd64; go build -o ./bin/$(BINARY_UNIX) ./cmd/$(BINARY_NAME).go

install:
	go build -o $(GOPATH)/bin/$(BINARY_NAME) main.go metadata.go migration.go

custom-install:
	go build -o $(INSTALL_PATH)/$(BINARY_NAME) main.go metadata.go migration.go

test:
	@echo "Not implemented"

grc-test:
	@echo "Not implemented"

clean:
	go clean
	rm -f ./bin/$(BINARY_NAME)
	rm -f ./bin/$(BINARY_UNIX)

## Misc
custom-build:
	make mod tidy; go mod vendor; go build ./...

get-deps:
	go get -u "github.com/davecgh/go-spew"
	go get -u "gopkg.in/yaml.v2"
