#!/bin/bash
seven apply \
  --config sevenconfig.yaml \
  --manifest 02-generate-data.yaml \
  --logs \
  --output data.md
