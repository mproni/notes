package checking

import (
	"fmt"
	"os"

	"github.com/mproni/notes/internal/handler"
)

func Arguments() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: note [OPTION] ...")
		return
	}

	if len(os.Args) > 5 {
		fmt.Println("Too many arguments. Use note -h (--help) for more info")
		return
	}

	if len(os.Args) == 2 {
		if os.Args[1] == "-rall" || os.Args[1] == "--read-all" {
			handler.ReadAll()
		} else if os.Args[1] == "-h" || os.Args[1] == "--help" {
			handler.Help()
		} else {
			fmt.Println("//some ERROR args 2")
		}
		return
	}

	if len(os.Args) == 3 {
		if os.Args[1] == "-r" || os.Args[1] == "--read" {
			handler.Read()
		} else if os.Args[1] == "-d" || os.Args[1] == "--delete" {
			handler.Delete()
		} else {
			fmt.Println("//some ERROR args 3")
		}
		return
	}

	if len(os.Args) == 4 {
		switch os.Args[1] {
		case "-a", "--add":
			handler.Add()
		case "-ut", "--update-title":
			handler.UpdateTitle()
		case "-uc", "--update-content":
			handler.UpdateContent()
		default:
			fmt.Println("//some ERROR args 4")
		}
		return
	}

	if len(os.Args) == 5 {
		if os.Args[1] == "-u" || os.Args[1] == "--update" {
			handler.Update()
		}
	}
}
