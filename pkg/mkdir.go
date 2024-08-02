package pkg

import (
	"fmt"
	"os"
)

func mkdir(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("mkdir: missing operand")
	}

	for _, dir := range args {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
