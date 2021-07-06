#!/usr/bin/env sh

#USAGE: ./lint.sh <folder_to_lint>

golint $1/... -set_exit_status

