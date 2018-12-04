#!/bin/bash

# See https://github.com/go-swagger/go-swagger
# install swagger codegen for
go get github.com/go-swagger/go-swagger

# Fetch
COMMERCE_STOREFRONT_URL="https://electronics.cqz1m-softwarea1-d10-public.model-t.cc.commerce.ondemand.com"
curl -o commerce-swagger-schema.json "$COMMERCE_STOREFRONT_URL/rest/v2/api-docs"

# Codegen swagger classes into the ./client folder
swagger generate client -f ./commerce-swagger-schema.json -A commercesuite

# Correct datetime format parsing
perl -pi -e 's/strfmt.DateTime/CustomDateTime/g' models/*.go
perl -pi -e 's/strfmt.DateTime/CustomDateTime/g' client/*.go