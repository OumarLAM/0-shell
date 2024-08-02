package pkg

import (
	"fmt"
	"io"
	"os"
)

func cp(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("cp: missing file operand")
	}

	src, dst := args[0], args[1]

	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}
