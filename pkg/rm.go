package pkg

import (
	"fmt"
	"os"
)

func removeFiles(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("\x1b[31mrm: missing operand\x1b[0m")
	}

	recursive := false
	files := []string{}

	// Check if the second argument is "-r"
	if len(args) > 0 && args[0] == "-r" {
		recursive = true
		// Append all files, skipping the second argument
		files = append(files, args[1:]...)
	} else {
		// All arguments are considered as files
		files = args
	}

	if len(files) == 0 {
		return fmt.Errorf("\x1b[31mrm: missing file operand\x1b[0m")
	}

	for _, file := range files {
		err := remove(file, recursive)
		if err != nil {
			return fmt.Errorf("\x1b[31mrm: %v\x1b[0m", err)
		}
	}

	return nil
}

func remove(path string, recursive bool) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("\x1b[31mrm: %v\x1b[0m", err)
	}

	if !info.IsDir() {
		err = os.Remove(path)
		if err != nil {
			return fmt.Errorf("\x1b[31mrm: %v\x1b[0m", err)
		}
		return nil
	}

	if !recursive {
		return fmt.Errorf("\x1b[31mrm: cannot remove '%s': Is a directory\x1b[0m", path)
	}

	err = os.RemoveAll(path)
	if err != nil {
		return fmt.Errorf("\x1b[31mrm: %v\x1b[0m", err)
	}
	return nil
}
