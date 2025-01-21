#!/bin/bash

# Путь к файлу, если он находится в репозитории
FILE="path_to_file_in_repo"

# Проверка существования файла
if [ -f "$FILE" ]; then
    # Вывод первых двух строк и последних двух строк
    (head -n 2 "$FILE"; echo ""; tail -n 2 "$FILE")
else
    echo "Ошибка: Файл $FILE не найден"
fi


