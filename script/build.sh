#!/bin/bash
go env -w CGO_ENABLED=0

servers=(
    "web"
)

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

if [[ "${servers[@]}"  =~ "${target}" ]]; then
    go build  -o main ./cmd/$target
    docker build -t $target:$version .
    rm ./main
elif [[ ! "${servers[@]}"  =~ "${target}" ]]; then
    echo "target $target not exists"
fi