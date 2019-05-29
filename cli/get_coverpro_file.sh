#!/usr/bin/env bash

set -e

go test ./... -v -covermode=count -coverprofile=coverage.out

bash <(curl -s https://codecov.io/bash) -t $CODECOV_TOKEN
goveralls -coverprofile=coverage.out -service=travis-ci -repotoken $COVERALLS_TOKEN
