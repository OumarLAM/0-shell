package pkg

import (
	"fmt"
	"strings"
)

func ExecuteCommand(input string) error {
	cmd, args := ParseCommand(input)

	switch cmd {
	case "echo":
		return Echo(args)
	case "cd":
		return Cd(args)
	case "ls":
		return Ls()
	case "pwd":
		return Pwd()
	case "cat":
		return Cat()
	case "cp":
		return Cp()
	case "rm":
		return Rm()
	case "mkdir":
		return Mkdir()
	default:
		return fmt.Errorf("command %s not found: ", cmd)
	}
}

func ParseCommand(input string) (string, []string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return "", nil
	}
	return parts[0], parts[1:]
}
