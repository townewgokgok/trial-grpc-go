#!/bin/sh

mkdir -p model
protoc -I proto proto/person.proto --go_out=plugins=grpc:model
