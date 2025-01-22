#!/bin/bash

# Найти путь к директории mystery
MYSTERY_DIR=$(find . -type d -name "mystery" | head -n 1)

# Проверить, найден ли путь
if [ -z "$MYSTERY_DIR" ]; then
    echo "Directory 'mystery' not found!"
    exit 1
fi

# Перейти в найденную директорию
cd "$MYSTERY_DIR"

# Найти убийцу
MURDERER=$(grep "CLUE" crimescene | grep "murderer" | awk -F':' '{print $2}' | xargs)
echo "The murderer is: $MURDERER"
