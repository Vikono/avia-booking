GOCMD=go
GOTEST=$(GOCMD) test
BINARY_NAME=build.out

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CIAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

all: build

build: 
	go build -o $(BINARY_NAME) ./cmd/app/

gen-api: 
	go generate ./...

lint: 
	golangci-lint run -c ./.golangci.yml

test:
	go test -v hello.go

run:
	go run -C ./cmd/app/ .

clean:
	go clean
	rm $(BINARY_NAME)