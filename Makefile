APP=kappa-web

.PHONY: build
## build: build kappa-web application
build:
	@echo "Building..."
	go build -o ${APP} main.go

.PHONY: run
## run: run main.go
run:
	go run -race main.go

.PHONY: clean
## clean: clean up the binary
clean:
	@echo "Cleaning..."
	go clean

.PHONY: help
## help: print the help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /' 
