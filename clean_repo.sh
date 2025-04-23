#!/bin/bash

# Переключаемся на ветку first-iteration
git checkout second-iteration

# Создаем новую чистую ветку без истории
git checkout --orphan temp-clean

# Добавляем все файлы (кроме игнорируемых)
git add -A

# Создаем финальный коммит
git commit -m "последний комит"

# Переключаемся на main
git checkout main

# Перезаписываем main новой чистой версией
git reset --hard temp-clean

# Принудительно обновляем удаленный main (осторожно!)
git push -f origin main

# Удаляем временные ветки
git branch -D temp-clean
git branch -D second-iteration
git push origin --delete second-iteration

echo "Готово! Репозиторий очищен, последний коммит: 'последний комит'"