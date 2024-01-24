#!/bin/bash

# go helpers

function pre_gen_go {
  # Create the new directory if it doesn't exist
  if [ ! -d $1 ]; then
    mkdir -p $1
  fi

  # Check if the go.mod exists
  if [ ! -f "$1/go.mod" ]; then
    # Create the file
    cd $1
    go mod init github.com/kiyani-org/proto-solution
    cd -
  fi
}

function post_gen_go() {
  cd $1
  go mod tidy
  cd -
}

# ts helpers

function pre_gen_ts {
  # Create the new directory if it doesn't exist
  if [ ! -d $1 ]; then
    mkdir -p $1
  fi

  # Check if package.json exists
  if [ ! -f "$1/package.json" ]; then
    cp perm/package.json $dest_root_dir
  fi
  cd $1
  npm i
  cd -
}

function main() {
  pre_gen_go gen/go
  # pre_gen_ts gen/ts

  buf generate
  if [ $? -ne 0 ]; then
    echo "Failed buf generate"
    exit 1
  fi

  post_gen_go gen/go
}

main 
