package pkg

import (
	"fmt"
	"strings"
)

func ExecuteCommand(input string) error {
	cmd, args := ParseCommand(input)

	switch cmd {
	case "echo":
		return echo(args)
	case "cd":
		return cd(args)
	case "ls":
		return ls(args)
	case "pwd":
		return pwd()
	case "cat":
		return cat(args)
	case "cp":
		return cp(args)
	case "rm":
		return rm(args)
	case "mv":
		return mv(args)
	case "mkdir":
		return mkdir(args)
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
