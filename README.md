Это простейшее CLI приложение для создания заметок.

Использование:
1. Добавить заметку:
   `./note.exe -a <title> <content>`
   `./note.exe --add <title> <content>`
2. Показать заметку с конкретным id:
   `./note.exe -r <id>`
   `./note.exe --read <id>`
3. Показать все заметки:
   `./note.exe -rall`
   `./note.exe --read-all`
4. Обновить заметку:
   `./note.exe -u <id> <title> <content>`
   `./note.exe --update <id> <title> <content>`
5. Обновить title заметки:
   `./note.exe -ut <id> <title>`
   `./note.exe --update-title <id> <title>`
6. Обновить content заметки:
   `./note.exe -uc <id> <content>`
   `./note.exe --update-content <id> <content>`
7. Удалить заметку с конкретным id:
   `./note.exe -d <id>`
   `./note.exe --delete <id>`
8. Помощь:
   `./note.exe -h`
   `./note.exe --help`