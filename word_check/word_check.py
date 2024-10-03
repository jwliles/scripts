#!/usr/bin/env python3

import string
import re

# Input and output file paths
input_file = "word_list.txt"
output_file = "clean_list.txt"

# Function to check if a word is English (only contains alphabetic characters A-Z)
def is_english_word(word):
    return re.match(r'^[a-zA-Z]+$', word) is not None

# Function to check if a word has at least 4 letters
def is_min_length(word, min_length=4):
    return len(word) >= min_length

# Function to check for three or more consecutive identical letters
def has_consecutive_repeats(word, limit=3):
    count = 1
    for i in range(1, len(word)):
        if word[i] == word[i - 1]:
            count += 1
            if count >= limit:
                return True
        else:
            count = 1
    return False

# Function to check for repeated letters at the start or end of the word
def has_repeated_edges(word):
    # Check if word starts with 2+ repeated letters
    if re.match(r'^([a-zA-Z])\1+', word):
        return True
    # Check if word ends with 2+ repeated letters
    if re.search(r'([a-zA-Z])\1+$', word):
        return True
    return False

# Set to store unique words
unique_words = set()

# Read and process the input file
with open(input_file, "r") as f:
    for line in f:
        # Strip punctuation and convert to lowercase
        word = line.strip().lower().translate(str.maketrans('', '', string.punctuation))

        # Apply all checks: only English characters, meets length, no consecutive repeats, no repeated edges
        if is_english_word(word) and is_min_length(word) and not has_consecutive_repeats(word) and not has_repeated_edges(word):
            unique_words.add(word)

# Write the cleaned words to the output file
with open(output_file, "w") as f:
    for word in sorted(unique_words):
        f.write(word + "\n")

print(f"Cleaned words have been saved to {output_file}")
