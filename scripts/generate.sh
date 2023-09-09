#!/bin/bash

# Recursively loop through the current directory and its subdirectories
function generate_protos_rec() {
  for file_or_dir in "$1"/*; do
    echo "entry $file_or_dir"
    trimmed_path=${file_or_dir#$raw_protos_dir/}

    if [ -f "$file_or_dir" ]; then
      echo "trimmed $trimmed_path"

      # Copy the file_or_dir to the new directory
      echo "copying ($file_or_dir) to ($dest_root_dir/$trimmed_path)"
      export PATH="$PATH:$(go env GOPATH)/bin" && protoc --go_out=$dest_root_dir $file_or_dir
    elif [ -d "$file_or_dir" ]; then
      echo "mkdir ($dest_root_dir/$trimmed_path)"
      # Recursively copy the contents of the subdirectory
      generate_protos_rec "$file_or_dir"
    fi
  done
}

function pre_gen_go {
  # Create the new directory if it doesn't exist
  if [ ! -d $1 ]; then
    mkdir -p $1
  fi

  # Check if the go.mod exists
  if [ ! -f "$1/go.mod" ]; then
    # Create the file
    echo "init go module"
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

function main() {
  # go protos
  dest_root_dir=gen/go
  pre_gen_go $dest_root_dir
  generate_protos_rec $raw_protos_dir
  post_gen_go $dest_root_dir
}

# Get the current directory
raw_protos_dir=protos

main 
