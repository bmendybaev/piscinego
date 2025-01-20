#!/bin/bash

# Проверяем, установлена ли переменная HERO_ID
if [ -z "$HERO_ID" ]; then
  echo "HERO_ID is not set. Please set it using: export HERO_ID=<id>"
  exit 1
fi

# URL JSON-данных
URL="https://01.tomorrow-school.ai/assets/superhero/all.json"

# Получаем relatives и применяем gsub для удаления символов новой строки
RELATIVES=$(curl -s "$URL" | jq -r --arg id "$HERO_ID" '.[] | select(.id == ($id | tonumber)) | .connections.relatives | gsub("\\n"; " ") | gsub("  +"; " ")')

# Проверяем, найдены ли данные
if [ -z "$RELATIVES" ]; then
  echo "No relatives found for HERO_ID=$HERO_ID."
  exit 1
fi

# Выводим результат
echo "$RELATIVES"
