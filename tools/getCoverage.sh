#!/bin/sh

cd kadai3-1/task4233/eimg
go test -coverprofile=profile ./...
go tool cover -html=profile
