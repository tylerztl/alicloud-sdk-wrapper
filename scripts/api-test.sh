#!/usr/bin/env bash

# test run instances API

curl -s -X POST \
  http://localhost:8080/cloud/run \
  -H "content-type: application/x-www-form-urlencoded" \
  -d ''