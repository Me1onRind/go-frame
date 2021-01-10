#!/bin/bash

cd `dirname $0`
protoc --proto_path=./:. --micro_out=./pb --go_out=./pb *.proto

pb_go_files=`ls -1 ./pb/ | grep -v 'micro.pb.go'`
for file in ${pb_go_files[@]}
do
    if ! command -v protoc-go-inject-tag > /dev/null 2>&1; then
        `cd ~ && go get github.com/favadi/protoc-go-inject-tag`
    fi
    protoc-go-inject-tag -input=./pb/${file}
done
