#!/bin/bash

# Directory containing your audio files
AUDIO_DIR="./audio"

# Temporary directory for converted files
TEMP_DIR="./converted/"

# Output directory for transcriptions
OUTPUT_DIR="./transcriptions/"

# Model path
MODEL_PATH="/home/jwl/projects/whisper.cpp/models/ggml-small.en-tdrz.bin"

# Path to the main executable
MAIN_EXEC="/home/jwl/projects/whisper.cpp/main"

# Create the temporary and output directories if they don't exist
mkdir -p "$TEMP_DIR"
mkdir -p "$OUTPUT_DIR"

# Count total files for progress tracking
TOTAL_FILES=$(ls -1 "$AUDIO_DIR"/* | wc -l)
CURRENT_FILE=0

# Convert all files first
echo "Starting conversion of $TOTAL_FILES files..."
for FILE in "$AUDIO_DIR"/*; do
  CURRENT_FILE=$((CURRENT_FILE + 1))
  # Get the file extension
  EXT="${FILE##*.}"
  # Get the base name of the file without extension
  BASENAME=$(basename "$FILE" .$EXT)
  # Convert the file to 16 kHz WAV and store it in the temporary directory
  ffmpeg -i "$FILE" -ar 16000 "$TEMP_DIR/$BASENAME.wav"
  echo "Converted $CURRENT_FILE of $TOTAL_FILES: $FILE"
done

# Transcribe all converted files
TOTAL_CONVERTED=$(ls -1 "$TEMP_DIR"/*.wav | wc -l)
CURRENT_CONVERTED=0

echo "Starting transcription of $TOTAL_CONVERTED files..."
for FILE in "$TEMP_DIR"/*.wav; do
  CURRENT_CONVERTED=$((CURRENT_CONVERTED + 1))
  # Get the base name of the file without extension
  BASENAME=$(basename "$FILE" .wav)
  # Transcribe the converted file using the tinydiarize option and save the output to a text file
  "$MAIN_EXEC" -m "$MODEL_PATH" -f "$FILE" --tinydiarize >"$OUTPUT_DIR/$BASENAME.txt"
  echo "Transcribed $CURRENT_CONVERTED of $TOTAL_CONVERTED: $FILE"
done

# Clean up the temporary files
rm -r "$TEMP_DIR"

echo "All files processed."
