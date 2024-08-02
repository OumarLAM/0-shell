package pkg

import (
	"fmt"
	"io"
	"os"
)

func cat(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("cat: missing file operand")
	}

	for _, filename := range args {
		file, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(os.Stdout, file)
		if err != nil {
			return err
		}
	}
	return nil
}
