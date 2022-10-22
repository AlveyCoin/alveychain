#!/bin/sh

set -e

ALVEY_CHAIN_BIN=./alveychain

case "$1" in

   "init")
      echo "Generating secrets..."
      secrets=$("$ALVEY_CHAIN_BIN" secrets init --num 4 --data-dir data- --json)
      echo "Secrets have been successfully generated"

      echo "Generating genesis file..."
      "$ALVEY_CHAIN_BIN" genesis \
        --dir /genesis/genesis.json \
        --consensus ibft \
        --ibft-validators-prefix-path data- \
        --bootnode /dns4/node-1/tcp/1478/p2p/$(echo $secrets | jq -r '.[0] | .node_id') \
        --bootnode /dns4/node-2/tcp/1478/p2p/$(echo $secrets | jq -r '.[1] | .node_id')
      echo "Genesis file has been successfully generated"
      ;;

   *)
      echo "Executing alveychain..."
      exec "$ALVEY_CHAIN_BIN" "$@"
      ;;

esac
