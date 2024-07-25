package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	if !(len(os.Args) == 4 || len(os.Args) == 2 && os.Args[1] == "all") {
		fmt.Println("Usage: note -[flag] [title] [content]")
		return
	}

	if os.Args[1] != "-add" && os.Args[1] != "all" {
		fmt.Println("Incorrect flag")
		return
	}

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

	if os.Args[1] == "-add" {
		statement.Exec(os.Args[2], os.Args[3])
	}

	if os.Args[1] == "all" {
		row, err := db.Query("SELECT id, title, content FROM notes ORDER BY id")
		if err != nil {
			log.Fatal(err)
		}
		defer row.Close()

		for row.Next() {
			var id int
			var title string
			var content string
			row.Scan(&id, &title, &content)
			fmt.Println("ID:", id, "Title:", title, "Content:", content)
		}
	}
}
