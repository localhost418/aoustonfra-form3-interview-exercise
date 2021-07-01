#!/usr/bin/env bash

scripts/wait-for/wait-for-http-code.sh 'http://accountapi:8080/v1/health' '200' && \
scripts/wait-for/wait-for-status.sh 'http://accountapi:8080/v1/health' 'up' && \
make integration_tests
