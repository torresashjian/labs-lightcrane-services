#!/bin/bash

network_type=${1:?}
os_type=${2:?}
arch_type=${3:?}

if [[ "${arch_type}" == "amd64" ]]; then
    pushd ./oss > /dev/null || exit 1
    ./stop.sh
    popd || exit 1
fi