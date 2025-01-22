#!/bin/bash

cd mystery
MURDERER=$(grep "CLUE" crimescene | grep "murderer" | awk -F':' '{print $2}' | xargs)
echo "The murderer is: $MURDERER"

