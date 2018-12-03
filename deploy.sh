#!/bin/bash

set -euxo

kubectl -n mini-stage1 apply -f deployment.yaml