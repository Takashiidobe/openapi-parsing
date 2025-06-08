#!/bin/bash

autorest \
  --input-file=../azure-rest-api-specs/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-08-15/cosmos-db.json \
  --v3 \
  --use:@autorest/modelerfour \
  --output-artifact=openapi-document \
  --modelerfour.lenient-model-deduplication=true \
	--clear-output-folder=true \
  --output-folder=./json/cosmos_db
