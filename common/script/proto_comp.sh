#!/bin/sh

protoc --go_out=common/email_service_pb/ --go_opt=paths=source_relative \
    --go-grpc_out=common/email_service_pb/ --go-grpc_opt=paths=source_relative \
    $1