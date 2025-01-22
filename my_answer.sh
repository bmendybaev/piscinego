#!/bin/bash

# Locate the mystery directory
MYSTERY_DIR=$(find . -type d -name "mystery" | head -n 1)

if [ -z "$MYSTERY_DIR" ]; then
    echo "Directory 'mystery' not found!"
    exit 1
fi

cd "$MYSTERY_DIR" || { echo "Failed to enter 'mystery' directory!"; exit 1; }

# Ensure the crimescene file exists
if [ ! -f crimescene ]; then
    echo "File 'crimescene' not found!"
    exit 1
fi

# Extract all lines containing the word 'CLUE' and save them to a temporary file
echo "Collecting clues from 'crimescene':"
grep "CLUE" crimescene > clues.txt

# Analyze the clues to find the murderer
MURDERER=$(grep -i "murderer" clues.txt | awk -F':' '{print $2}' | xargs)

if [ -z "$MURDERER" ]; then
    echo "No murderer identified in the clues!"
else
    echo -e "\nThe murderer is: $MURDERER"
fi

# Clean up the temporary file
rm -f clues.txt

# If hints are available, print them
for hint in hint1 hint2 hint3; do
    if [ -f "$hint" ]; then
        echo -e "\nHint from '$hint':"
        cat "$hint"
    fi
done
