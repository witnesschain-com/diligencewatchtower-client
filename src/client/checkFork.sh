#!/bin/bash

# RPCs
OP_GETH="http://35.153.33.57:8545/"
SEQUENCER="https://goerli.optimism.io"

NOW=$(cast block-number --rpc-url $OP_GETH)
BLOCK_DIFF=$(diff <(cast block $NOW --rpc-url $OP_GETH) <(cast block $NOW --rpc-url $SEQUENCER))

# Check if BLOCK_DIFF is empty
if [[ -z "$BLOCK_DIFF" ]]; then
        echo "No diff in block #$NOW between op-geth and canonical RPC."
else
        echo "-----------------------------------------------------------------------"
        echo " -> ERROR! Diff in block #$NOW between op-geth and canonical RPC"
        echo "$BLOCK_DIFF"
        echo "-----------------------------------------------------------------------"
fi