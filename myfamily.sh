#!/bin/bash

if [ -z "$HERO_ID" ]; then
  echo "HERO_ID is not set. Please set it using: export HERO_ID=<id>"
  exit 1
fi

URL="https://01.tomorrow-school.ai/assets/superhero/all.json"

RELATIVES=$(curl -s "$URL" | jq -r --arg id "$HERO_ID" '.[] | select(.id == ($id | tonumber)) | .connections.relatives' | tr '\n' ' ' | sed 's/  */ /g' | sed 's/^ *//;s/ *$//')

if [ -z "$RELATIVES" ]; then
  echo "No relatives found for HERO_ID=$HERO_ID."
  exit 1
fi

echo "$RELATIVES"
