#!/usr/bin/env bash

cmdname=$(basename $0)

function usage()
{
    cat << USAGE >&2
Usage:
    $cmdname <url> <http_code> [timeout]
USAGE
    exit 1
}


if [ "$#" -lt 2 ]; then
    usage
fi

URL=$1
HTTP_CODE=$2
TIMEOUT=$3
TIMEOUT=${TIMEOUT:-15}

function get_http_code(){
    echo $(curl -s -o /dev/null -w "%{http_code}" $1)
}

function wait_for() {
    echo "waiting for ${URL} http code ${HTTP_CODE} for ${TIMEOUT} seconds"
    while true; do
        current_http_code=$(get_http_code ${URL})
        if [ ! -z "${current_http_code}" ]  && [ ${current_http_code} -eq ${HTTP_CODE} ]; then
            echo "http_code for ${URL} is ${HTTP_CODE}"
            exit 0
        fi
        sleep 1
        if [ ${SECONDS} -ge ${TIMEOUT} ]; then
            exit 2
        fi
    done
}

wait_for
