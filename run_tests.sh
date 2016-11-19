#!/bin/sh

#apt-get update
#apt-get install -y inotify-tools

RunTests () {
  docker-compose -f docker-compose.test.yml -p jnotes-test run api \
    go test -v $(go list ./... | grep -v /vendor/) \
    | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
}

RunTests
while true; do
  inotifywait -e modify,create,delete -r --exclude '(gin-bin|.git)' . && RunTests
done
