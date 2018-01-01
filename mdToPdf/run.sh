#!/bin/sh
echo =============================================== Testing
go run mdToPdf.go -i test.md -o test.pdf -debug=true > debug.txt