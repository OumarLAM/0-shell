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
		if len(args) < 2 {
			return fmt.Errorf("cd: missing operand")
		}
		return os.Chdir(args[1])
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("pwd: %v", err)
		}
		fmt.Println(dir)
		return nil
	case "ls":
		// Implémentation simplifiée de ls
		// return listDirectory(args[1:])
		return nil
	case "cat":
		return nil

		// return concatenateFiles(args[1:])
	case "cp":
		return nil
		// return copyFiles(args[1:])
	case "rm":
		return nil
		// return removeFiles(args[1:])
	case "mv":
		return nil
		// return moveFiles(args[1:])
	case "mkdir":
		if len(args) < 2 {
			return fmt.Errorf("mkdir: missing operand")
		}
		return os.Mkdir(args[1], 0755)
	case "exit":
		os.Exit(0)
		return nil
	default:
		return fmt.Errorf("\x1b[31mcommand `%s` not found\x1b[0m", args[0])
	}
}
