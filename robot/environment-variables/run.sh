#!/bin/bash
export AI_NAME="Seven of Nine"
export NAME="Jean-Luc Picard"
seven apply \
  --config sevenconfig.yaml \
  --manifest use-env-vars.yaml
