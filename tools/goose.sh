#! /bin/sh

set -a
. ../.env
set +a

goose -dir $GOOSE_MIGRATION_DIR $GOOSE_DRIVER $GOOSE_DBSTRING $1