#!/bin/bash

function pre_gen_go {
  echo "running pre_gen_go"
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

function post_gen_go {
  echo "running post_gen_go"
  cd $1
  go mod tidy
  cd -
}

function pre_gen_es {
  echo "running pre_gen_es"

  # Create the new directory if it doesn't exist
  if [ ! -d $1 ]; then
    mkdir -p $1
  fi

  # Check if package.json exists
  if [ ! -f "$1/package.json" ]; then
    cp ./perm/es/package.json $1
  fi

  if [ -f "$1/index.ts" ]; then
    touch $1/index.ts
  fi
}

function post_gen_es {
  echo "running post_gen_es"
  cd $1
  npm i
  cd -
}

function main {
  pre_gen_go gen/go
  pre_gen_es gen/es

  buf generate
  if [ $? -ne 0 ]; then
    echo "Failed buf generate"
    exit 1
  fi

  post_gen_go gen/go
  post_gen_es gen/es
}

main 
