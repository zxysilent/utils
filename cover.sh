#!/bin/bash

# 覆盖率测试

go test -v -coverprofile cover.out .
go tool cover -html=cover.out -o cover.html

echo "cover end"