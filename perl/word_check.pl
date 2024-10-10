#!/usr/bin/env perl

use strict;
use warnings;
use utf8;

# Input and output file paths
my $input_file = 'word_list.txt';
my $output_file = 'clean_word_list.txt';

# Hash to store unique words
my %unique_words;

# Open the input file for reading
open(my $in, '<:encoding(UTF-8)', $input_file) or die "Could not open '$input_file' $!";

# Process the file line-by-line
while (my $line = <$in>) {
    chomp $line;
    $line = lc($line); # Convert to lowercase
    $line =~ s/[[:punct:]]//g; # Remove punctuation

    # Check if word contains only English letters and has at least 4 characters
    next unless ($line =~ /^[a-z]{4,}$/);

    # Skip if there are three or more consecutive identical letters
    next if ($line =~ /(.)\1\1/);

    # Skip if the word starts or ends with two or more identical letters
    next if ($line =~ /^(.)(\1)+/ || $line =~ /(.)(\1)+$/);

    # Add to hash for uniqueness
    $unique_words{$line} = 1;
}

close($in);

# Open the output file for writing
open(my $out, '>:encoding(UTF-8)', $output_file) or die "Could not open '$output_file' $!";

# Write sorted unique words to output
foreach my $word (sort keys %unique_words) {
    print $out "$word\n";
}

close($out);

print "Cleaned words have been saved to $output_file\n";
