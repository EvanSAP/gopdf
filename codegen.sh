#!/bin/bash

# See https://github.com/go-swagger/go-swagger

swagger generate client -f ./commerce-swagger-schema.json -A commercesuite
