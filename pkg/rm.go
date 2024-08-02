package pkg

import (
	"fmt"
	"os"
)

func rm(args []string) error {
	var recursive bool
	var paths []string

	for _, arg := range args {
		if arg == "-r" {
			recursive = true
		} else {
			paths = append(paths, arg)
		}
	}

	if len(paths) == 0 {
		return fmt.Errorf("rm: missing operand")
	}

	for _, path := range paths {
		var err error
		if recursive {
			err = os.RemoveAll(path)
		} else {
			err = os.Remove(path)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
