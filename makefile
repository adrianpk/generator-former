# Vars
PROD_TAG=v0.0.1

# Misc
BINARY_NAME=mw
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build

build:
	go build -o ./bin/$(BINARY_NAME) main.go metadata.go migration.go

build-linux:
	CGOENABLED=0 GOOS=linux GOARCH=amd64; go build -o ./bin/$(BINARY_UNIX) ./cmd/$(BINARY_NAME).go

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
	go get -u "github.com/go-chi/chi"
	go get -u "github.com/jmoiron/sqlx"
	go get -u "github.com/kr/pretty"
	go get -u "github.com/lib/pq"
	go get -u "github.com/mattn/go-sqlite3"
	go get -u "github.com/satori/go.uuid"
	go get -u "gitlab.com/mikrowezel/backend/config"
	go get -u "gitlab.com/mikrowezel/backend/db"
	go get -u "gitlab.com/mikrowezel/backend/db/postgres"
	go get -u "gitlab.com/mikrowezel/backend/log"
	go get -u "gitlab.com/mikrowezel/backend/migration"
	go get -u "gitlab.com/mikrowezel/backend/service"
	go get -u "golang.org/x/crypto"
	go get -u "golang.org/x/net"
	go get -u "gopkg.in/check.v1"

