package handler

import (
	"fmt"

	"github.com/mproni/notes/internal/database"
)

func Add() {
	database.AddSQL()
	fmt.Println("Заметка добавлена")
}

func Read() {
	database.ReadSQL()
}

func ReadAll() {
	database.ReadAllSQL()
}

func Update() {
	database.UpdateSQL()
	fmt.Println("Заметка обновлена")
}

func UpdateTitle() {
	database.UpdateTitleSQL()
	fmt.Println("Заголовок заметки обновлен")
}

func UpdateContent() {
	database.UpdateContentSQL()
	fmt.Println("Содержимое заметки обновлено")
}

func Delete() {
	database.DeleteSQL()
	fmt.Println("Заметка удалена")
}

func Help() {
	fmt.Println("no help")
}
