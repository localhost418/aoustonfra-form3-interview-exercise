#!/usr/bin/env bash

cmdname=$(basename $0)

function usage()
{
    cat << USAGE >&2
Usage:
    $cmdname <url> <status> [timeout]
USAGE
    exit 1
}


if [ "$#" -lt 2 ]; then
    usage
fi

URL=$1
STATUS=$2
TIMEOUT=$3
TIMEOUT=${TIMEOUT:-15}

function get_health_status(){
    echo $(curl -s $1 | jq -r .status)
}

function wait_for() {
    echo "waiting for ${URL} status ${STATUS} for ${TIMEOUT} seconds"
    while true; do
        current_health_status=$(get_health_status ${URL})
        if [ ! -z "${current_health_status}" ]  && [ "${current_health_status}" = ${STATUS} ]; then
	    echo "${URL} status is ${STATUS}"
            exit 0
        fi
        sleep 1
	if [ ${SECONDS} -ge ${TIMEOUT} ]; then
            exit 2
        fi
    done
}

wait_for
