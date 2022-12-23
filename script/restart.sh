#!/bin/bash

home=${HOME}
wkdir=$(cd $(dirname $0); pwd)
cd $wkdir/..

target=$1
version=$2

if [ ! -n "$target" ] ;then
    echo "missing target"
    exit 1
fi
if  [ ! -n "$version" ];then
    echo "missing version"
    exit 1
fi

# 删除旧容器
name=${target}_${version}
echo "container name: $name"
container=$(docker ps -a  | grep  $name | awk '{print $1}')
if [ -n "$container" ]; then
    echo "docker rm -f $container"
    docker rm -f $container
fi

# 启动新容器
if [ $target = "web" ];then
    echo "docker run --rm -d -p 8080:8080 --name $name $target:$version"
    docker run --rm -d -p 8080:8080 --name $name \
        -v /etc/localtime:/etc/localtime \
        -v /etc/timezone:/etc/timezone \
        -v $home/logs:/apps/logs \
        -v $home/cfg:/apps/.cfg $target:$version
fi 