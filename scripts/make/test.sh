#!/usr/bin/env sh

#USAGE: ./test.sh <folder_with_test_files> [opts]

test_folder=$1
shift

go test ./${test_folder}/... $@
