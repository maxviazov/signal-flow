#!/bin/bash

echo "=== Go Environment Fix Script ==="
echo ""

echo "Current problematic environment:"
echo "GOPROXY: $GOPROXY"
echo "GOPRIVATE: $GOPRIVATE"
echo ""

echo "Setting correct environment variables for this session..."
export GOPROXY="https://proxy.golang.org,direct"
export GOPRIVATE="github.com/maxviazov/*"

echo "New environment:"
echo "GOPROXY: $GOPROXY"
echo "GOPRIVATE: $GOPRIVATE"
echo ""

echo "Testing go get commands..."
echo "Running: go get -v -u github.com/go-playground/validator/v10 github.com/spf13/viper"
echo ""

go get -v -u github.com/go-playground/validator/v10 github.com/spf13/viper

if [ $? -eq 0 ]; then
    echo ""
    echo "✅ SUCCESS: go get commands completed without errors!"
    echo ""
    echo "=== SOLUTION SUMMARY ==="
    echo "The issue was caused by malformed GOPROXY environment variable."
    echo "Fixed by separating GOPROXY and GOPRIVATE into distinct variables:"
    echo "  GOPROXY=\"https://proxy.golang.org,direct\""
    echo "  GOPRIVATE=\"github.com/maxviazov/*\""
    echo ""
    echo "=== PERMANENT FIX RECOMMENDATIONS ==="
    echo "1. Restart your terminal/shell to reload environment variables"
    echo "2. Restart GoLand/IntelliJ IDEA to pick up the corrected environment"
    echo "3. The workspace.xml file has been fixed for IDE-specific settings"
    echo "4. Your .zshrc appears to have correct settings - ensure no other scripts override them"
else
    echo ""
    echo "❌ ERROR: go get commands still failing"
    echo "Additional troubleshooting may be needed"
    exit 1
fi