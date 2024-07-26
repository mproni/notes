package handler

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/mproni/notes/internal/database"
)

func HandleIt(db *sql.DB) {
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "-rall", "--read-all":
			database.ReadAllSQL(db)
		case "-h", "--help":
			fmt.Println("No help, you have to deal with this on your own")
		}
		return
	}

	if len(os.Args) == 3 {
		switch os.Args[1] {
		case "-r", "--read":
			database.ReadSQL(db)
		case "-d", "--delete":
			database.DeleteSQL(db)
		default:
			fmt.Println("Incorrect flag")
		}
		return
	}

	if len(os.Args) == 4 {
		switch os.Args[1] {
		case "-a", "--add":
			database.AddSQL(db)
		case "-ut", "--update-title":
			database.UpdateTitleSQL(db)
		case "-uc", "--update-content":
			database.UpdateContentSQL(db)
		default:
			fmt.Println("Incorrect flag")
		}
		return
	}

	if len(os.Args) == 5 {
		switch os.Args[1] {
		case "-u", "--update":
			database.UpdateSQL(db)
		default:
			fmt.Println("Incorrect flag")
		}
		return
	}
}
