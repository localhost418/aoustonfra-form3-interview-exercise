#!/usr/bin/env sh

#USAGE: ./lint.sh <folder_to_lint>

output=$(golint $1/...)

if [ ! -z "$output" ]; then
  echo $output
  exit 1
fi
