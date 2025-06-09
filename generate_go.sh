#!/bin/bash

autorest \
  --input-file=../azure-rest-api-specs/specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/storage.json  \
  --v3 \
  --use:@autorest/modelerfour \
  --modelerfour.lenient-model-deduplication=true \
	--clear-output-folder=true \
  --output-folder=./go/storage \
	--go
