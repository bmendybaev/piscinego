#!/bin/bash

# Удаляем предыдущую копию репозитория, если она существует
rm -rf the-final-cl-test

# Клонируем репозиторий
git clone https://github.com/01-edu/the-final-cl-test.git the-final-cl-test

# Проверяем, успешно ли клонирован репозиторий
if [ -d "the-final-cl-test" ]; then
  # Если папка существует, выводим содержимое
  ls -R the-final-cl-test
else
  # Если папка не создана, выводим сообщение об ошибке
  echo "Ошибка: репозиторий не клонирован."
fi
