#!/bin/sh

GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o pkg/tropo-recording-catcher.exe recording-catcher.go
echo "Built for Windows..."
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o pkg/tropo-recording-catcher.linux recording-catcher.go
echo "Built for Linux..."
GOOS=darwin GOARCH=amd64  CGO_ENABLED=0 go build -o pkg/tropo-recording-catcher.osx recording-catcher.go
echo "Built for OSX..."