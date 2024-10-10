# Audio Conversion and Transcription Script

## Overview

This Bash script automates the process of converting audio files to 16 kHz WAV format and then transcribing them using a
speech recognition model. It uses `ffmpeg` for audio conversion and `whisper.cpp` for transcription. The transcriptions
are saved as text files for easy access.

## Features

- **Audio Conversion**: Converts all audio files in the specified directory to 16 kHz WAV format, stored temporarily.
- **Transcription**: Transcribes the converted WAV files using a specified model and saves the output in text format.
- **Progress Tracking**: Displays progress updates during both conversion and transcription steps.
- **Clean-up**: Automatically removes the temporary directory used for converted files after processing.

## Usage

1. **Set up directories and paths**:
    - **AUDIO_DIR**: Directory containing your original audio files (`./audio` by default).
    - **TEMP_DIR**: Temporary directory to store converted WAV files (`./converted/` by default).
    - **OUTPUT_DIR**: Directory to store transcription output (`./transcriptions/` by default).
    - **MODEL_PATH**: Path to the speech recognition model file (`ggml-small.en-tdrz.bin`).
    - **MAIN_EXEC**: Path to the main executable for transcription (`whisper.cpp`).

2. **Run the script**:
   ```bash
   ./your_script_name.sh
   ```

3. The script will:
    - Convert all audio files in `AUDIO_DIR` to WAV format.
    - Transcribe each WAV file using `whisper.cpp` with the `--tinydiarize` option.
    - Save the transcriptions as `.txt` files in `OUTPUT_DIR`.
    - Clean up the temporary converted files after processing.

## Requirements

- **Bash**: The script is designed to be run in a Bash shell.
- **ffmpeg**: Required for converting audio files to the desired format (16 kHz WAV).
- **whisper.cpp**: A speech recognition tool for transcribing audio. Ensure the paths to the model and main executable
  are correct.
- **Audio Files**: The script processes all files in the `AUDIO_DIR` directory, so ensure it contains supported audio
  formats (e.g., `.mp3`, `.wav`, etc.).

## Customization

- **Audio Directory**: Modify the `AUDIO_DIR` variable to point to your desired directory containing audio files.
- **Model and Executable Path**: Adjust `MODEL_PATH` and `MAIN_EXEC` to match your environment's setup.
- **Output Locations**: The output transcriptions will be saved in `OUTPUT_DIR`, which can be changed as needed.

## Example

```bash
./audio-to-text.sh
```

This will convert all audio files in `./audio/`, transcribe them, and save the transcriptions to `./transcriptions/`.

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.