#!/bin/bash

function execute_protoc() {
  dest_root_dir=$1
  file_or_dir=$2
  lang=$3
  case $lang in
    go)
      export PATH="$PATH:$(go env GOPATH)/bin" && protoc --go_out=$dest_root_dir $file_or_dir 
      ;;
    ts)
      protoc --plugin=$dest_root_dir/node_modules/.bin/protoc-gen-ts_proto \
             --ts_proto_out=$dest_root_dir \
             --ts_proto_opt=outputServices=grpc-js \
             --ts_proto_opt=esModuleInterop=true \
             $file_or_dir 
      ;;
  esac
}

# Recursively loop through the current directory and its subdirectories
function generate_protos_rec() {
  for file_or_dir in "$1"/*; do
    trimmed_path=${file_or_dir#$raw_protos_dir/}
    if [ -f "$file_or_dir" ]; then
      # Copy the file_or_dir to the new directory
      execute_protoc $dest_root_dir $file_or_dir $2
    elif [ -d "$file_or_dir" ]; then
      # Recursively copy the contents of the subdirectory
      generate_protos_rec "$file_or_dir" $2
    fi
  done
}

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
  # go protos
  dest_root_dir=gen/go
  pre_gen_go $dest_root_dir
  generate_protos_rec $raw_protos_dir "go"
  post_gen_go $dest_root_dir

  # ts protos
  dest_root_dir=gen/ts
  pre_gen_ts $dest_root_dir
  generate_protos_rec $raw_protos_dir "ts"
}

# Set the protos directory
raw_protos_dir=protos

main 
