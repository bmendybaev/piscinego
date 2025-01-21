#!/bin/bash

# URL репозитория
REPO_URL="https://github.com/01-edu/the-final-cl-test.git"

# Клонирование репозитория
git clone $REPO_URL

# Проверка содержимого репозитория
REPO_NAME="the-final-cl-test"
echo "Полное содержимое репозитория:"
ls -la $REPO_NAME
