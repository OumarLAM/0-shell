package pkg

import (
	"fmt"
	"os"
	"strings"

	"github.com/OumarLAM/0-shell/utils"
)

func ExecuteCommand(args []string, input string) error {
	switch args[0] {
	case "echo":
		fmt.Println(strings.TrimPrefix(input, "echo "))
		return nil
	case "cd":
		return changeDirectory(args[1:])
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			return utils.FormatError("pwd: %v", err)
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
			return utils.FormatError("mkdir: missing operand")
		}
		return os.Mkdir(args[1], 0755)
	case "chmod":
		if len(args) < 3 {
			return utils.FormatError("chmod: missing operand")
		}
		return ParseAndApplyChmod(args[1], args[2:])
	case "touch":
		if len(args) < 2 {
			return utils.FormatError("touch: missing file operand")
		}
		for _, filename := range args[1:] {
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

		return utils.FormatError("command `%s` not found", args[0])
	}
}
