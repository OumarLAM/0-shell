package pkg

import (
	"fmt"
	"os"
)

func moveFiles(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("\x1b[31mmv: missing file operand\x1b[0m")
	}

	source := args[0]
	destination := args[1]

	// Move or rename the file or directory
	err := os.Rename(source, destination)
	if err != nil {
		return fmt.Errorf("\x1b[31mmv: %v\x1b[0m", err)
	}

	return nil
}
