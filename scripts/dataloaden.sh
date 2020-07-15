#!/usr/bin/env bash

set -euox pipefail

SCRIPT_FILE=$(basename $0)
echo ${SCRIPT_FILE}

SCRIPT_DIR=$(dirname $0)
echo ${SCRIPT_DIR}

cd ${SCRIPT_DIR} && cd ../src/graph/model/

# https://github.com/vektah/dataloaden
go get -u github.com/vektah/dataloaden@0.3.0
dataloaden ContentLoader string []*github.com/sky0621/wolf-bff/src/adapter/controller/gqlmodel.Content
