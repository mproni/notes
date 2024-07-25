package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTableSQL() {
	db, err := sql.Open("sqlite3", "./notes.db")
	if err != nil {
		log.Fatal(err)
	}
	// ну ты и дурак, конечно.
	//	defer db.Close()

	createTableSQL := `CREATE TABLE IF NOT EXISTS notes (
		        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		        "title" TEXT,
		        "content" TEXT
		    );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}
