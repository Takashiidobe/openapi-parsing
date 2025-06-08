#!/bin/bash

autorest \
  --input-file=../azure-rest-api-specs/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-08-15/cosmos-db.json \
  --v3 \
  --use:@autorest/modelerfour \
  --modelerfour.lenient-model-deduplication=true \
	--clear-output-folder=true \
  --output-folder=./go/cosmos_db \
	--go
