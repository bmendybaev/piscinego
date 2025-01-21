#!/bin/bash

# 1. Клонируем репозиторий
git clone https://github.com/01-edu/the-final-cl-test.git the-final-cl-test >/dev/null 2>&1

# 2. Проверяем, существует ли файл final.txt
if [ -f "the-final-cl-test/final.txt" ]; then
  # 3. Извлекаем вторую строку из файла и выводим её
  head -2 the-final-cl-test/final.txt | tail -1
else
  echo "Ошибка: файл final.txt не найден."
fi
