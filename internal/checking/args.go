package checking

import (
	"fmt"
	"os"
)

func Arguments() bool {
	switch len(os.Args) {
	case 1:
		fmt.Println("Usage: note [OPTION] ...")
		return false
	case 2, 3, 4, 5:
		return keywords()
	default:
		fmt.Println("Too many arguments. Use note -h (--help) for more info")
		return false
	}
}

func keywords() bool {
	switch os.Args[1] {
	case "-h", "--help", "-rall", "--read-all", "-r", "--read",
		"-d", "--delete", "-a", "--add", "-ut", "--update-title",
		"-uc", "--update-content", "-u", "--update":
		return true
	default:
		return false
	}
}
