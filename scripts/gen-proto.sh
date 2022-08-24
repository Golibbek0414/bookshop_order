#!/bin/bash

CURRENT_DIR=$(pwd)

mkdir $CURRENT_DIR/genproto/catalog_service
mkdir $CURRENT_DIR/genproto/order_service

protoc -I /usr/local/include \
       -I $GOPATH/src/github.com/gogo/protobuf/gogoproto \
       -I $CURRENT_DIR/protos/catalog_service/ \
        --gofast_out=plugins=grpc:$CURRENT_DIR/genproto/catalog_service/ \
        $CURRENT_DIR/protos/catalog_service/*.proto;

protoc -I /usr/local/include \
       -I $GOPATH/src/github.com/gogo/protobuf/gogoproto \
       -I $CURRENT_DIR/protos/order_service/ \
        --gofast_out=plugins=grpc:$CURRENT_DIR/genproto/order_service/ \
        $CURRENT_DIR/protos/order_service/*.proto;

if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i "" -e "s/,omitempty//g" $CURRENT_DIR/genproto/*.go
  else
    sed -i -e "s/,omitempty//g" $CURRENT_DIR/genproto/*.go
fi