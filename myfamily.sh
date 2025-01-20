#!/bin/bash

# Check if HERO_ID is set
if [ -z "$HERO_ID" ]; then
  echo "HERO_ID is not set. Please set it using: export HERO_ID=<id>"
  exit 1
fi

# URL of the JSON file
URL="https://01.tomorrow-school.ai/assets/superhero/all.json"

# Fetch the data and filter by HERO_ID to get relatives
RELATIVES=$(curl -s "$URL" | jq -r --arg id "$HERO_ID" '.[] | select(.id == ($id | tonumber)) | .connections.relatives')

# Check if relatives were found
if [ -z "$RELATIVES" ]; then
  echo "No relatives found for HERO_ID=$HERO_ID."
  exit 1
fi

# Output the relatives without quotes
echo "$RELATIVES"
