#!/bin/bash

CURRENT_DIR="$1"

echo "Current directory: ${CURRENT_DIR}"

rm -rf "${CURRENT_DIR}/genproto"
echo "Removed existing genproto directory."

mkdir -p "${CURRENT_DIR}/genproto"
echo "Created new genproto directory."

echo "Processing proto files..."
protoc -I="${CURRENT_DIR}/protos" \
  --go_out="${CURRENT_DIR}/genproto" --go-grpc_out="${CURRENT_DIR}/genproto" \
  "${CURRENT_DIR}/protos"/*.proto

echo "Protobuf files have been generated."
