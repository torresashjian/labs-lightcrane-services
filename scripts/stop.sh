#!/bin/bash

network_type=${1:?}
os_type=${2:?}
arch_type=${3:?}
component_type=${4:?}

stop_backend(){
    pushd ./${arch_type}/lc-backend > /dev/null || exit 1
    ./stop.sh
    popd || exit 1
}

stop_agent(){
    pushd ./${arch_type}/lc-agent > /dev/null || exit 1
    ./stop.sh
    popd || exit 1
}

if [[ "${component_type}" == "backend" ]]; then
    stop_backend
elif [[ "${component_type}" == "agent" ]]; then
    stop_agent
fi