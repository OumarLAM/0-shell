package pkg

import (
	"fmt"
	"os"
	"strings"
)

func ExecuteCommand(args []string) error {
	switch args[0] {
	case "echo":
		fmt.Println(strings.Join(args[1:], " "))
		return nil
	case "cd":
		return changeDirectory(args[1:])
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("pwd: %v", err)
		}
		fmt.Println(dir)
		return nil
	case "ls":
		return listDirectory(args[1:])
	case "cat":
		return concatenateFiles(args[1:])
	case "cp":
		return copyFiles(args[1:])
	case "rm":
		return removeFiles(args[1:])
	case "mv":
		return moveFiles(args[1:])
	case "mkdir":
		if len(args) < 2 {
			return fmt.Errorf("mkdir: missing operand")
		}
		return os.Mkdir(args[1], 0755)
	case "exit":
		os.Exit(0)
		return nil
	case "clear":
		clear()
		return nil
	default:
		return fmt.Errorf("\x1b[31mcommand `%s` not found\x1b[0m", args[0])
	}
}
