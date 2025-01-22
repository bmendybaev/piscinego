#!/bin/bash

cd mystery

# Найти все строки с CLUE и извлечь любые упоминания о murder
grep "CLUE" crimescene | grep -i "murder" | awk -F':' '{print $2}' | xargs > potential_murderer.txt

if [ -s potential_murderer.txt ]; then
    MURDERER=$(cat potential_murderer.txt)
    echo "The murderer is: $MURDERER"
else
    echo "DDDDDCould not identify the murderer from the clues!"
fi

# Очистка временного файла
rm -f potential_murderer.txt
