#!/usr/bin/env python3

import i3ipc


def split_dynamic(i3):
    focused = i3.get_tree().find_focused()
    if not focused:
        return

    parent = focused.parent
    if not parent:
        return

    # Check if the focused container is in a tabbed or stacked layout
    if parent.layout in ["tabbed", "stacked"]:
        # Navigate up to the parent container of the tabbed or stacked layout
        parent = parent.parent

    # Decide split direction based on the container dimensions
    if parent.rect.height > parent.rect.width:
        i3.command("split h")
    else:
        i3.command("split v")


if __name__ == "__main__":
    i3 = i3ipc.Connection()
    split_dynamic(i3)
