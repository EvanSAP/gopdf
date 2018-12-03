#!/bin/bash

set -euxo

# kubectl -n mini-stage1 delete -f deployment.yaml
kubectl -n mini-stage1 apply -f deployment.yaml