#!/bin/bash

# Проверяем, установлена ли переменная HERO_ID
if [ -z "$HERO_ID" ]; then
  echo "HERO_ID is not set. Please set it using: export HERO_ID=<id>"
  exit 1
fi

# URL с данными о супергероях
URL="https://01.tomorrow-school.ai/assets/superhero/all.json"

# Получаем данные и обрабатываем relatives
RELATIVES=$(curl -s "$URL" | jq -r --arg id "$HERO_ID" '.[] | select(.id == ($id | tonumber)) | .connections.relatives | gsub("\\n"; ", ")')

# Проверяем, найдены ли родственники
if [ -z "$RELATIVES" ]; then
  echo "No relatives found for HERO_ID=$HERO_ID."
  exit 1
fi

# Выводим родственников
echo "$RELATIVES"
