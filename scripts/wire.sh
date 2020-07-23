#!/usr/bin/env bash

set -euox pipefail

SCRIPT_FILE=$(basename $0)
echo ${SCRIPT_FILE}

SCRIPT_DIR=$(dirname $0)
echo ${SCRIPT_DIR}

cd ${SCRIPT_DIR} && cd ../src/setup

# https://github.com/google/wire
go get -u github.com/google/wire/cmd/wire@v0.4.0
wire
