#!/usr/bin/env bash

# test run instances API

curl -X POST "http://localhost:8080/v1/cloud/run" \
  -H  "accept: application/json" \
  -H "content-type: application/json" \
  -d "{}"