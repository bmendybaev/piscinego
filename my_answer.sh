#!/bin/bash

# Перейти в директорию mystery
cd mystery || { echo "Directory 'mystery' not found!"; exit 1; }

# Проверить наличие файла crimescene
if [ ! -f crimescene ]; then
    echo "File 'crimescene' not found in 'mystery' directory!"
    exit 1
fi

# Найти и вывести все строки с "CLUE" в crimescene
echo "Searching for clues in 'crimescene'..."
grep "CLUE" crimescene

# Если файл hint
