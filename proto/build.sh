#!/bin/bash

cd `dirname $0`
protoc --proto_path=./:. --micro_out=./pb --go_out=./pb *.proto

pb_go_files=`ls -1 ./pb/ | grep -v 'micro.pb.go'`
for file in ${pb_go_files[@]}
do
    protoc-go-inject-tag -input=./pb/${file}
done
