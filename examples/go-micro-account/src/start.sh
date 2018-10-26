#!/usr/bin/env bash
CONSUL=127.0.0.1:8500

if [ $# == 1 ] && [ $1 == "dev" ]
    then
        echo "start develop mode"
        echo "复制配置文件"
        cp ../config/dev.json ./env.json

        read -p "输入consul ip:" -t 30 ip
        if [ $ip ] 
            then
                echo "当前不为空为空，设置consul-->$ip" 
                CONSUL=$ip
        else
            echo "当前ip为空，使用默认值$CONSUL"
        fi
        echo "当前注册中心ip---->$CONSUL"
        CONSUL_HTTP_ADDR=$CONSUL go run ./*.go  --registry=consul --registry_address=$CONSUL --selector=cache --server=grpc --client=grpc
fi

if [ $# == 1 ]
    then
        echo "start $1 mode"
        echo "当前注册中心ip---->$CONSUL"
        echo "复制配置文件"

        cp "../config/$1.json" "./env.json"

        go build ./*.go

        CONSUL_HTTP_ADDR=$CONSUL ./app --registry=consul --registry_address=$CONSUL --selector=cache --server=grpc --client=grpc
fi