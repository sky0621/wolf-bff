#!/usr/bin/env bash

set -euox pipefail

SCRIPT_FILE=$(basename $0)
echo ${SCRIPT_FILE}

SCRIPT_DIR=$(dirname $0)
echo ${SCRIPT_DIR}

cd ${SCRIPT_DIR} && cd ../src/

rm -fr ./adapter/controller/dummyresolvers/*.go
rm -fr ./adapter/controller/graphql_generated.go
rm -fr ./adapter/controller/graphql_model_gen.go

# https://gqlgen.com/
# https://github.com/99designs/gqlgen
go get -u github.com/99designs/gqlgen@v0.11.3
gqlgen
