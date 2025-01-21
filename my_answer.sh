#!/bin/bash

# Клонирование репозитория
REPO_URL="https://github.com/01-edu/the-final-cl-test.git"
TARGET_DIR="the-final-cl-test"

# Проверка, существует ли директория
if [ ! -d "$TARGET_DIR" ]; then
    git clone "$REPO_URL"
fi

# Печать строки
echo "John Doe$"
