#!/bin/sh

cd kadai2/task4233/eimg
go test -coverprofile=profile ./...
go tool cover -html=profile
