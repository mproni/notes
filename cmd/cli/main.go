package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./notes.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTableSQL := `CREATE TABLE IF NOT EXISTS notes (
		        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		        "title" TEXT,
		        "content" TEXT
		    );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	insertNoteSQL := `INSERT INTO notes (title, content) VALUES (?, ?)`
	statement, err := db.Prepare(insertNoteSQL)
	if err != nil {
		log.Fatal(err)
	}
	// всё не работает, потому что нужно CGO для компиляции.
	fmt.Println(len(os.Args), os.Args[0], os.Args[1])
}
