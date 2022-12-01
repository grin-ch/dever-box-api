#!/bin/bash

servers=(
    "web"
)

pid=$$
target=$1

if [ $target != "" ];then
    if [[ "${servers[@]}"  =~ "${target}" ]]; then
        echo "run.sh pid:${pid}"
        echo "target:${target} in serve"
        go run ./cmd/$target
    elif [[ ! "${servers[@]}"  =~ "${target}" ]]; then
        echo "target $target not exists"
    fi
else
    echo "missing arg"
fi