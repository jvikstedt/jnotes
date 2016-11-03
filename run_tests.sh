#!/bin/sh

go test $(go list ./... | grep -v /vendor/)
while true; do
  inotifywait -e modify,create,delete -r . && go test $(go list ./... | grep -v /vendor/)
done
