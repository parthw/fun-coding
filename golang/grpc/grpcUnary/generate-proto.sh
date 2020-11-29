#!/bin/bash

#protoc --proto_path=pb greet.proto --go_out=plugins=grpc:.
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/greet.proto