#!/bin/bash

curl -o documentation.json https://devnet-dec2-explorer-api.decimalchain.com/api/documentation.json
./gpretty.py
./swagger2code.py
gofmt -w endpoints.go
mv endpoints.go ../api
