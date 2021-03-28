APP=kappa-web
PKG_LIST := $(shell go list ${APP}/... | grep -v /vendor/)

.PHONY: lint
## Lint Golang files
lint:
	@golint -set_exit_status ${PKG_LIST}

.PHONY: build
## build: build kappa-web application
build:
	@echo "Building..."
	@go build -mod vendor -o ${APP} main.go

.PHONY: run
## run: run main.go
run:
	@go run -mod vendor -race main.go

.PHONY: clean
## clean: clean up the binary
clean:
	@echo "Cleaning..."
	@go clean

.PHONY: sync
## sync: sync go vendor modules
sync:
	@echo "Syncing go modules..."
	@go mod tidy \
		&& go mod vendor

.PHONY: test
## test: test kappa-web functions
test:
	@echo "Testing..."
	@go test -v -race -cover=true .

.PHONY: test-coverage
## Run tests with coverage
test-coverage:
	@go test -v -race -coverprofile cover.out -covermode=atomic .
	@cat cover.out >> coverage.txt

.PHONY: help
## help: print the help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /' 
