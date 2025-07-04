PROJECTNAME=$(shell basename $(PWD))
GOFILES = $(wildcard */*.go)

# Make is verbose in Linux. Made it silent
MAKEFLAGS += --silent

## usage: make [option]
## Options:
.PHONY : help
help : Makefile
	echo "\n$(PROJECTNAME) is an example custom blockchain built in golang\n"
	sed -n 's/^## //p' $<
	sed -n 's/^###//p' $<

### build: Build the project
.PHONY : build
build:
	@echo "  >  Building binary..."
	@go build -o bin/$(PROJECTNAME)
	@echo "  >  Binary build complete - bin/$(PROJECTNAME)"

### run: Run the project
.PHONY : run
run: build
	@echo "  >  Running binary..."
	@./bin/$(PROJECTNAME)

### test: Run tests
.PHONY : test
test:
	@echo "  >  Running tests..."
	@go test -v ./...

### proto: Generate protobuf files
.PHONY : proto
proto:
	@echo "  >  Generating protobuf files..."
	@protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	 ./proto/*.proto
	@echo "  >  Protobuf files generated"