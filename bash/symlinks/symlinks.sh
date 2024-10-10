#!/bin/bash

# Set the source directory (dotfiles/.emacs.d) and the target directory (~/.emacs.d)
SOURCE_DIR=~/dotfiles/.emacs.d
TARGET_DIR=~/.emacs.d

# Create the main .emacs.d directory if it doesn't exist
mkdir -p $TARGET_DIR

# List of files and directories to symlink
FILES=()

# Create symlinks
for FILE in "${FILES[@]}"; do
  ln -s $SOURCE_DIR/$FILE $TARGET_DIR/$FILE
done
