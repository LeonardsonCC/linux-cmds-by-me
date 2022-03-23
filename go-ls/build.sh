#! /bin/bash
GOOS=linux go build -ldflags "-s -w" -o build/go-ls
