#!/bin/bash

SOURCE_FILE="proxy.go"
OUTPUT_LINUX="proxy-linux"
OUTPUT_WINDOWS="proxy-windows.exe"

echo "Compiling for Linux..."
GOOS=linux GOARCH=amd64 go build -o $OUTPUT_LINUX $SOURCE_FILE

echo "Compiling for Windows..."
GOOS=windows GOARCH=amd64 go build -o $OUTPUT_WINDOWS $SOURCE_FILE

echo "Compilation completed!"
