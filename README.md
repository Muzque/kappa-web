# Kappa Web
The Kappa Web is for calculating your crypto coins profits!
Functions:
- Accounting your every wallet.
- Calculate the profit of every kind of coins.
- Draw your incomes and drawback diagrams.

## Table of Contents
- [Setup](#setup)
    - [On Mac OS](#on-macos)
- [Build and Install](#build-and-install)
- [Run Unit Tests](#run-unit-tests)


## Setup

### On macOS
Install Go and set environment variables `GOPATH`, `GOBIN`, and `PATH`. The current code base should compile with **Go 1.16.2**. On macOS, install Go with the following command
```
brew install go@1.16.2
brew link go@1.16.2 --force
```

Clone this repo into your `$GOPATH`. The path should look like this: `$GOPATH/src/github.com/kappa-web/`

```
git clone https://github.com/Muzque/kappa-web.git $GOPATH/src/github.com/kappa-web/
export KAPPA_HOME=$GOPATH/src/github.com/kappa-web/
cd $KAPPA_HOME
```

## Build and Install
This should build the binary and copy them into your $GOPATH/bin. One binary kappa-web will be generated. kappa-web can be regarded as the launcher of the Kappa web service.

```
export GO111MODULE=on
make build
```

## Run Unit Tests
Run unit tests with the command below

```
make test
```
