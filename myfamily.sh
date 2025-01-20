#!/bin/bash

# Проверяем, задана ли переменная HERO_ID
if [ -z "$HERO_ID" ]; then
  echo "HERO_ID is not set. Please set the environment variable HERO_ID."
  exit 1
fi

# Загружаем JSON с данными о супергероях
DATA=$(curl -s https://01.tomorrow-school.ai/assets/superhero/all.json)

# Извлекаем информацию о родственниках субъекта по ID
RELATIVES=$(echo "$DATA" | jq -r --arg id "$HERO_ID" '.[] | select(.id == ($id | tonumber)) | .connections.relatives')

# Проверяем, найдены ли данные
if [ -z "$RELATIVES" ]; then
  echo "No relatives found for HERO_ID=$HERO_ID."
  exit 1
fi

# Выводим результат
echo "$RELATIVES"
