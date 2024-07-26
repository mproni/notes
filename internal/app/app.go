package app

import (
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/mproni/notes/internal/checking"
	"github.com/mproni/notes/internal/database"
	"github.com/mproni/notes/internal/handler"
)

func Start() {
	if !(checking.Arguments()) {
		os.Exit(1)
	}

	db := database.CreateTableSQL()
	defer db.Close()

	handler.HandleIt(db)
}
