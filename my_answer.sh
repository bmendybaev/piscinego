#!/bin/bash

# Клонирование репозитория
git clone https://github.com/01-edu/the-final-cl-test.git the-final-cl-test >/dev/null 2>&1

# Проверка, успешно ли клонирован репозиторий
if [ ! -d "the-final-cl-test" ]; then
  echo "Ошибка: репозиторий не клонировался."
  exit 1
fi

# Проверка наличия файла final.txt
if [ ! -f "the-final-cl-test/final.txt" ]; then
  echo "Ошибка: файл final.txt не найден."
  ls -R the-final-cl-test  # Выводим структуру папки для отладки
  exit 1
fi

# Извлечение второй строки
cat the-final-cl-test/final.txt | head -2 | tail -1
