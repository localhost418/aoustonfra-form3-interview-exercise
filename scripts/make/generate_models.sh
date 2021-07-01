#!/usr/bin/env sh

#USAGE: ./generate_models <swagger_spec_file> <output_folder>

rm -rf $2
mkdir  $2

# generate model for Account ressource
swagger generate model --with-flatten expand --skip-validation -f $1 -t $2 -n Account
