#!/bin/bash

set -e
cd $(dirname $(pwd))/cmd

function copy_config_exec() {
    cd ..
    PLUGIN_HD_PATH=$(find /usr/lib -type d -name dde-grand-search-daemon)
    sudo cp cmd/translate-search-plugin ${PLUGIN_HD_PATH}/plugins/searcher/
    sudo cp config/translate-search-plugin.conf ${PLUGIN_HD_PATH}/plugins/searcher/
}

function build_amd64() {
    export GOOS=linux
    export GOARCH=amd64
    go build -o translate-search-plugin .
    copy_config_exec
}
function build_loong64() {
    export GOOS=linux
    export GOARCH=loong64
    go build -o translate-search-plugin .
    copy_config_exec
}

function build_arm() {
    export GOOS=linux
    export GOARCH=arm
    go build -o translate-search-plugin .
    copy_config_exec
}

which go
if [[ $? -ne 0 ]]; then
    echo "未检查到 golang 环境"
    exit 1
fi

export GOPROXY='https://goproxy.cn,direct'

ARCH=$(uname -m)

case ${ARCH} in
x86_64)
    build_amd64
    ;;
loong64)
    build_loong64
    ;;
arm)
    build_arm
    ;;
esac
