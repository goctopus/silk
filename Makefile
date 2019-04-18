# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

all: run

test:
	$(GOTEST) -v ./

fmt:
	$(GOCLEAN)
	go fmt .
	go fmt ./dialect/
	go fmt ./drivers/mssql/
	go fmt ./drivers/mysql/
	go fmt ./drivers/postgres/
	go fmt ./drivers/sqlite/
	go fmt ./example/
	go fmt ./example/models/
	go fmt ./connection/

