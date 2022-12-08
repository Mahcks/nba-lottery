#!/bin/bash

release() {
  gox -osarch="windows/amd64"
  gox -osarch="darwin/amd64"
  gox -osarch="linux/amd64"
}

"$@"