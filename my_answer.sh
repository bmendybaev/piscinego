#!/bin/bash

# Navigate to the mystery directory
cd mystery || { echo "Directory 'mystery' not found!"; exit 1; }

# Ensure the crimescene file exists
if [ ! -f crimescene ]; then
    echo "File 'crimescene' not found!"
    exit 1
fi

# Extract all lines containing the word 'CLUE'
echo "Collecting clues from 'crimescene':"
grep "CLUE" crimescene

# If hints are available, print them
for hint in hint1 hint2 hint3; do
    if [ -f "$hint" ]; then
        echo -e "\nHint from '$hint':"
        cat "$hint"
    fi
done
