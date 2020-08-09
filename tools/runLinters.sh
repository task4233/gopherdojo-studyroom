#!/bin/sh

cd kadai2/task4233/eimg
echo "run goimports"
goimports -w .

echo "run gofmt"
gofmt -w .

echo "run golint"
golint ./...

echo "run gsc"
gsc ./...

echo "run gosec"
gosec ./...

echo "run staticcheck"
staticcheck ./...
