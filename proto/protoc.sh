#!/bin/sh

set -xe

SERVER_OUTPUT_DIR=/output/server
CLIENT_OUTPUT_DIR=/output/client

protoc --version
protoc -I=/proto/protos/chat chat.proto\
  --go_out=plugins="grpc:${SERVER_OUTPUT_DIR}" \
  --dart_out="grpc:${CLIENT_OUTPUT_DIR}"

protoc -I=/proto/protos/google/protobuf timestamp.proto wrappers.proto empty.proto\
  --dart_out="grpc:${CLIENT_OUTPUT_DIR}/google/protobuf"