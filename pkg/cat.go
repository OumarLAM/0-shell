package pkg

import (
	"fmt"
	"os"
)

func concatenateFiles(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("\x1b[31mcat: missing file operand\x1b[0m")
	}

	for _, filename := range args {
		data, err := os.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("\x1b[31mcat: %v\x1b[0m", err)
		}
		fmt.Print(string(data))
	}

	return nil
}
