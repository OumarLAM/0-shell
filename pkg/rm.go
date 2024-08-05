package pkg

import (
	"os"
	"strings"

	"github.com/OumarLAM/0-shell/utils"
)

func removeFiles(args []string) error {
	if len(args) == 0 {
		return utils.FormatError("rm: missing operand")
	}

	recursive := false
	files := []string{}

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			if strings.Contains(arg, "r") {
				recursive = true
			}
		} else {
			files = append(files, arg)
		}
	}

	if len(files) == 0 {
		return utils.FormatError("rm: missing file operand")
	}

	for _, file := range files {
		err := remove(file, recursive)
		if err != nil {
			return utils.FormatError("rm: %v", err)
		}
	}

	return nil
}

func remove(path string, recursive bool) error {
	info, err := os.Stat(path)
	if err != nil {
		return utils.FormatError("rm: %v", err)	}

	if !info.IsDir() {
		err = os.Remove(path)
		if err != nil {
			return utils.FormatError("rm: %v", err)		}
		return nil
	}

	if !recursive {
		return utils.FormatError("rm: cannot remove '%s': Is a directory", path)	}

	err = os.RemoveAll(path)
	if err != nil {
		return utils.FormatError("rm: %v", err)	}
	return nil
}
