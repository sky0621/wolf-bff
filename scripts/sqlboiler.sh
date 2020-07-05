#!/usr/bin/env bash

set -euox pipefail

SCRIPT_FILE=$(basename $0)
echo ${SCRIPT_FILE}

SCRIPT_DIR=$(dirname $0)
echo ${SCRIPT_DIR}

cd ${SCRIPT_DIR} && cd ../src/

# https://github.com/volatiletech/sqlboiler
go get -u github.com/volatiletech/sqlboiler/v4@v4.2.0
go get -u github.com/volatiletech/null/v8
sqlboiler --wipe psql
