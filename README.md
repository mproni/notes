Это простейшее CLI приложение для создания заметок.

Использование:
1. Добавить заметку:

   `./note.exe -a <title> <content>`

   `./note.exe --add <title> <content>`
3. Показать заметку с конкретным id:

   `./note.exe -r <id>`

   `./note.exe --read <id>`
4. Показать все заметки:

   `./note.exe -rall`

   `./note.exe --read-all`
5. Обновить заметку:

   `./note.exe -u <id> <title> <content>`

   `./note.exe --update <id> <title> <content>`
6. Обновить title заметки:

   `./note.exe -ut <id> <title>`

   `./note.exe --update-title <id> <title>`
7. Обновить content заметки:

   `./note.exe -uc <id> <content>`

   `./note.exe --update-content <id> <content>`
8. Удалить заметку с конкретным id:

   `./note.exe -d <id>`

   `./note.exe --delete <id>`
9. Помощь:

   `./note.exe -h`

   `./note.exe --help`
