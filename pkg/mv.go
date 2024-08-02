package pkg

import (
	"fmt"
	"os"
)

func mv(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("mv: missing file operand")
	}

	src, dst := args[0], args[1]
	return os.Rename(src, dst)
}
