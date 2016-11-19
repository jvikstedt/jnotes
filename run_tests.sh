#!/bin/sh

apt-get update
apt-get install -y inotify-tools

go test $(go list ./... | grep -v /vendor/)
while true; do
  inotifywait -e modify,create,delete -r --exclude '(gin-bin|.git)' . && go test $(go list ./... | grep -v /vendor/)
done
