#!/bin/bash

# Get the current directory
current_dir=$1

# Get the name of the new directory
new_dir_name=$2

protos_prefix=protos/

# Create the new directory if it doesn't exist
if [ ! -d $new_dir_name ]; then
  mkdir -p $new_dir_name
fi

# Recursively loop through the current directory and its subdirectories
function copy_files() {
  for file_or_dir in "$1"/*; do
    echo "entry $file_or_dir"
    trimmed_path=${file_or_dir#$protos_prefix}

    if [ -f "$file_or_dir" ]; then
      echo "trimmed $trimmed_path"

      # Copy the file_or_dir to the new directory
      echo "copying ($file_or_dir) to ($new_dir_name/$trimmed_path)"
      cp "$file_or_dir" "$new_dir_name/$trimmed_path"
    elif [ -d "$file_or_dir" ]; then
      echo "mkdir ($new_dir_name/$trimmed_path)"

      mkdir "$new_dir_name/$trimmed_path"
      # Recursively copy the contents of the subdirectory
      copy_files "$file_or_dir"
    fi
  done
}

# Copy the files from the current directory to the new directory
copy_files $current_dir
