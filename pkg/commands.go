package pkg

import (
	"os"
	"strings"

	"github.com/OumarLAM/0-shell/utils"
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
	case "chmod":
		if len(arguments) < 3 {
			return utils.FormatError("chmod: missing operand")
		}
		return ParseAndApplyChmod(arguments[1], arguments[2:])
	case "touch":
		if len(arguments) < 2 {
			return utils.FormatError("touch: missing file operand")
		}
		for _, filename := range arguments[1:] {
			if err := touchFile(filename); err != nil {
				return utils.FormatError("touch: %v", err)
			}
		}
		return nil
	case "exit":
		os.Exit(0)
		return nil
	case "clear":
		clear()
		return nil
	default:
		return utils.FormatError("command `%s` not found", arguments[0])
	}
}

func ParseCommand(input string) (string, []string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return "", nil
	}
	return parts[0], parts[1:]
}
