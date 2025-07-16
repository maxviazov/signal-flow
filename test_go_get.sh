#!/bin/bash

echo "Testing Go environment setup after fixing workspace.xml..."
echo "Current Go environment:"
go env GOPROXY
go env GOPRIVATE

echo ""
echo "Testing go get commands that were previously failing..."
echo "Running: go get -v -u github.com/go-playground/validator/v10 github.com/spf13/viper"

# Run the same command that was failing in the issue
go get -v -u github.com/go-playground/validator/v10 github.com/spf13/viper

if [ $? -eq 0 ]; then
    echo ""
    echo "SUCCESS: go get commands completed without errors!"
else
    echo ""
    echo "ERROR: go get commands still failing"
    exit 1
fi