#!/usr/bin/env bash
npm run build
go build -o ./cmd/cmd ./cmd
./cmd/cmd
