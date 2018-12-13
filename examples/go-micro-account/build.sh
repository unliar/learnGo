#!/bin/bash
set -e
env=("prod" "dev" "beta")
envStatus=0

for i in "${env[@]}"
do  
    if [[ "$i" == "$1" ]];then
       envStatus=1
       echo "current env $1 start to build"      
    fi
done

if [[ $envStatus -eq 0 ]];then
   echo "no matched env exit"
   exit 1
fi

# 程序名称
FileNameBuildOutPut=main
# 工作目录
WorkDir=app
# 配置目录
ConfigDir=config

go version

http_proxy=http://127.0.0.1:1080 dep ensure -v

go build -o $FileNameBuildOutPut app.go

echo "i wanna copy config $1"

rm -rf $WorkDir

mkdir $WorkDir && mkdir $WorkDir/$ConfigDir

echo "copy $1 config file starting"

cp "config/$1.env.json" "app/config/env.json"

cp $FileNameBuildOutPut $WorkDir

rm $FileNameBuildOutPut

echo "copy $1 config file done"

echo "build successful and copy config file"