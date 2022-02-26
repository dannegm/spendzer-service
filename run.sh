#!/usr/bin/env bash
go build -o spendzer main.go
codesign -s - spendzer
