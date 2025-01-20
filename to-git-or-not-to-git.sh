#!/bin/bash

# URL JSON данных
URL="https://01.tomorrow-school.ai/assets/superhero/all.json"

# Идентификатор супергероя
SUPERHERO_ID=170

# Получение данных (имя, сила и пол)
curl -s $URL | jq -r ".[] | select(.id == $SUPERHERO_ID) | .name, .powerstats.power, .appearance.gender"
