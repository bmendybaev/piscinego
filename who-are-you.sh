#!/bin/bash

# Define the URL where the JSON data is stored
URL="https://01.tomorrow-school.ai/assets/superhero/all.json"

# Define the Subject ID
SUBJECT_ID=70

# Fetch the name from the JSON data
curl -s $URL | jq -r ".[] | select(.id == $SUBJECT_ID) | .name" | awk '{print "\"" $0 "\""}'
