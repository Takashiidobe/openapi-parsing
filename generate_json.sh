#!/bin/bash

autorest \
  --input-file=../azure-rest-api-specs/specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/storage.json \
  --v3 \
  --use:@autorest/modelerfour \
  --output-artifact=openapi-document \
  --modelerfour.lenient-model-deduplication=true \
	--clear-output-folder=true \
  --output-folder=./json/storage
