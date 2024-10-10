#!/bin/bash

# Use the current directory as DIR
DIR="$(pwd)"

# Total Size in Gigabytes
echo "---------------------------------------"
echo "Directory Information for: $DIR"
echo "---------------------------------------"
echo "Total Size (in GB):"
du -sh "$DIR" | awk '{print $1}'

# Total Number of Files
echo "---------------------------------------"
echo "Total Files:"
find "$DIR" -type f | wc -l

# Total Number of Folders
echo "---------------------------------------"
echo "Total Folders:"
find "$DIR" -type d | wc -l

# Total Number of Empty Files
echo "---------------------------------------"
echo "Total Empty Files:"
find "$DIR" -type f -empty | wc -l

# Total Number of Empty Directories
echo "---------------------------------------"
echo "Total Empty Directories:"
find "$DIR" -type d -empty | wc -l
echo "---------------------------------------"
