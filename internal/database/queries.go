package database

import (
	"fmt"
	"log"
	"os"
)

func AddSQL() {
	query := `INSERT INTO notes (title, content) VALUES (?, ?)`
	_, err := db.Exec(query, os.Args[2], os.Args[3])
	if err != nil {
		log.Fatal(err)
	}
}

func ReadSQL() {
	var id int
	var title string
	var content string
	query := `SELECT id, title, content FROM notes WHERE id = ?`
	err := db.QueryRow(query, os.Args[2]).Scan(&id, &title, &content)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ID:", id, "Title:", title, "Content:", content)
}

func ReadAllSQL() {
	query := `SELECT id, title, content FROM notes ORDER BY id`
	row, err := db.Query(query)
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

func UpdateSQL() {
	query := `UPDATE notes SET title = ?, content = ? WHERE id = ?`
	_, err := db.Exec(query, os.Args[3], os.Args[4], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateTitleSQL() {
	query := `UPDATE notes SET title = ? WHERE id = ?`
	_, err := db.Exec(query, os.Args[3], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateContentSQL() {
	query := `UPDATE notes SET content = ? WHERE id = ?`
	_, err := db.Exec(query, os.Args[3], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteSQL() {
	query := `DELETE FROM notes WHERE id = ?`
	_, err := db.Exec(query, os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
}
