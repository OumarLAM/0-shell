package pkg

import (
	"fmt"
	"os"
	"strings"
)

func ExecuteCommand(arguments []string, input string) error {
	cmd, args := ParseCommand(input)

	switch cmd {
	case "echo":
		return echo(args)
	case "cd":
		return changeDirectory(args)
	case "ls":
		return listDirectory(args)
	case "pwd":
		return printWorkingDirectory()
	case "cat":
		return concatenateFiles(args)
	case "cp":
		return copyFiles(args)
	case "rm":
		return removeFiles(args)
	case "mv":
		return moveFiles(args)
	case "mkdir":
		return makeDirectory(args)
	case "exit":
		os.Exit(0)
		return nil
	case "clear":
		clear()
		return nil
	default:
		return fmt.Errorf("\x1b[31mcommand `%s` not found\x1b[0m", cmd)
	}
}

func ParseCommand(input string) (string, []string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return "", nil
	}
	return parts[0], parts[1:]
}
