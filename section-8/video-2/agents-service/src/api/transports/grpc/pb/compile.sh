#!/usr/bin/env sh

protoc agents-service.proto --go_out=plugins=grpc:.