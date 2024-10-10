#!/bin/bash
# update_crontab.sh
CRONTAB_FILE="/home/jwl/dotfiles/.scripts/cronjobs"
if [ -f "$CRONTAB_FILE" ]; then
    crontab $CRONTAB_FILE
else
    echo "Crontab file does not exist: $CRONTAB_FILE"
    exit 1
fi
