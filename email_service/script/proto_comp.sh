#!/bin/sh

protoc --go_out=email_service_pb --go_opt=paths=source_relative \
    --go-grpc_out=email_service_pb --go-grpc_opt=paths=source_relative \
    $1