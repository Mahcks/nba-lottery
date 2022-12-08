#!/bin/bash

# Build binaries for release
build() {
  gox -osarch="windows/amd64" -output="./build/nba-lottery-windows-amd64"
  gox -osarch="darwin/amd64" -output="./build/nba-lottery-darwin-amd64"
  gox -osarch="linux/amd64" -output="./build/nba-lottery-linux-amd64"
}

# Remove all files from `build` than re-build
re-build() {
  rm -rfv build/*
  build
}

# Builds a test from current source
test() {
  rm -rfv test/*
  gox -osarch="windows/amd64" -output="./test/nba-lottery-alpha"
}

"$@"