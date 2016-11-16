#!/bin/sh

go run schemes/build_tl_scheme.go < schemes/api_layer_55.json > tl_schema55.go
gofmt -w tl_schema.go
