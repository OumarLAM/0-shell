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
	case "pwd":
		return Pwd()
	// TODO: Add other commands
	default:
		return fmt.Errorf("command not found: %s", cmd)
	}
}

func ParseCommand(input string) (string, []string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return "", nil
	}
	return parts[0], parts[1:]
}
