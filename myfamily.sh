#!/bin/bash

# URL of the JSON file
URL="https://01.tomorrow-school.ai/assets/superhero/all.json"

# Check if HERO_ID is set
if [ -z "$HERO_ID" ]; then
  echo "Error: HERO_ID environment variable is not set."
  exit 1
fi

# Fetch superhero name based on HERO_ID and remove quotes
curl -s $URL | jq -r ".[] | select(.id == $HERO_ID) | .name"
