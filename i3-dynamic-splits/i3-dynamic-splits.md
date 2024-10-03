# Dynamic i3 Window Splitter

## Overview

This Python script dynamically adjusts the window split direction in the i3 window manager based on the dimensions of
the focused container. It uses the `i3ipc` library to interact with the i3 tree structure, determining whether to split
a window horizontally or vertically for optimal layout.

## Features

- **Automatic Split Direction**: Determines the split direction based on the dimensions of the parent container:
    - **Horizontal Split** (`split h`): If the height of the container is greater than its width.
    - **Vertical Split** (`split v`): If the width of the container is greater than its height.
- **Tab/Stack Layout Navigation**: If the focused window is within a tabbed or stacked layout, the script navigates up
  to the appropriate parent container for determining the split direction.

## Usage

1. **Install the `i3ipc` Python Package**: This script requires the `i3ipc` library to interact with the i3 window
   manager. Install it via `pip`:
   ```bash
   pip install i3ipc
   ```

2. **Run the Script**: Execute the script to apply the dynamic split:
   ```bash
   python3 split_dynamic.py
   ```

   You can bind this script to a keyboard shortcut in your i3 config file (`~/.config/i3/config`):
   ```plaintext
   bindsym $mod+s exec python3 /path/to/split_dynamic.py
   ```

   This allows you to trigger the script using a key combination (e.g., `$mod+s`).

## Requirements

- **Python 3**: The script is written in Python 3.
- **i3 Window Manager**: The script is designed to work with the i3 window manager, leveraging its tree layout for
  window management.
- **`i3ipc` Library**: Required to interact with i3's IPC (Inter-Process Communication) interface.

## Customization

- **Modify Split Logic**: You can adjust the logic that determines the split direction based on other conditions by
  editing the `split_dynamic` function.
- **Keyboard Shortcut**: Customize the key binding in your i3 config to trigger this script as desired.

## Notes

- **Parent Container Check**: The script automatically navigates up to the parent container if the focused window is
  part of a tabbed or stacked layout.
- **Script Location**: Make sure to provide the correct path to the script in your i3 config file when setting up a
  keyboard shortcut.

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.