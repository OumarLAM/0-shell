package pkg

import (
	"fmt"
	"os"
)

func makeDirectory(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("\x1b[31mmkdir: missing operand\x1b[")
	}

	for _, dir := range args {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
