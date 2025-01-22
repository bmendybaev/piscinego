#!/bin/bash

cd "$(dirname "$0")/mystery"
MURDERER=$(grep "CLUE" crimescene | grep "murderer" | awk -F':' '{print $2}' | xargs)
echo "The murderer is: $MURDERER"

