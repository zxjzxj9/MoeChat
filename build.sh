#! /bin/bash

GO_SRV_DIR=src/server_go
cd $GO_SRV_DIR
export GOPATH=`pwd`:$GOPATH
make
#make test
cd - 
