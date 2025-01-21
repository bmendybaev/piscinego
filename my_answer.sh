#!/bin/bash

# URL репозитория
REPO_URL="https://github.com/01-edu/the-final-cl-test"

# Клонирование репозитория
git clone $REPO_URL

# Печать содержимого репозитория
REPO_NAME="the-final-cl-test"
echo "Содержимое репозитория:"
ls -la $REPO_NAME
