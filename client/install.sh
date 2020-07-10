#!/bin/bash

UNAME=$(uname)
if [[ "${UNAME}" == CYGWIN* || "${UNAME}" == MINGW* || "${UNAME}" == MSYS* ]]; then
  # If using Git Bash on Windows, it will attempt to transform paths to windows Paths. Setting MSYS_NO_PATHCONV stops this process
  export MSYS_NO_PATHCONV=1
fi

docker run --rm -v .:/client -w /client node:12-slim bash -c "npm install"