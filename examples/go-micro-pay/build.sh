#!/bin/bash

# 程序名称
FileNameBuildOutPut=main
# 工作目录
WorkDir=app
# 配置目录
ConfigDir=config

go env

http_proxy=http://127.0.0.1:1080 dep ensure -v

go build -o $FileNameBuildOutPut *.go

echo "i wanna go config $1"

rm -rf $WorkDir

mkdir $WorkDir && mkdir $WorkDir/$ConfigDir

echo "copy $1 config file starting"

cp "config/$1.env.json" "app/config/env.json"

cp $FileNameBuildOutPut $WorkDir

rm $FileNameBuildOutPut

echo "copy $1 config file done"

echo "time to start you app in folder workdir app"