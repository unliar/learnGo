#!/usr/bin/env bash
CONSUL=127.0.0.1:8500

if [ $# == 1 ] && [ $1 == "dev" ]
    then
        echo "start develop mode"
        read -p "输入consul ip:" -t 30 ip
        if [ $ip ] 
            then
                echo "当前不为空为空，设置consul-->$ip" 
                CONSUL=$ip
        else
            echo "当前ip为空，使用默认值$CONSUL"
        fi
        echo "当前注册中心ip---->$CONSUL"
        CONSUL_HTTP_ADDR=$CONSUL go run app.go  --registry=consul --registry_address=$CONSUL --selector=cache --server=grpc --client=grpc
else
    echo "start production mode"
    CONSUL=$1
    echo "当前注册中心ip---->$CONSUL"
    go build app.go
    CONSUL_HTTP_ADDR=$CONSUL    ./app --registry=consul --registry_address=$CONSUL --selector=cache --server=grpc --client=grpc
fi


