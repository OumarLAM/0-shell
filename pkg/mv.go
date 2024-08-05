package pkg

import (
	"os"
	"path/filepath"

	"github.com/OumarLAM/0-shell/utils"
)

func moveFiles(args []string) error {
	if len(args) < 2 {
		return utils.FormatError("not enough arguments to move")
	}

	destination := args[len(args)-1]
	destinationInfo, err := os.Stat(destination)
	// Check if the destination exists and is a directory
	isDestinationDir := false
	if err == nil {
		if destinationInfo.IsDir() {
			isDestinationDir = true
		}
	} else if !os.IsNotExist(err) {
		return utils.FormatError("error accessing destination: %v", err)
	} else if filepath.Ext(destination) == "" {
		// If the destination does not exist and has no extension, treat as a directory
		err = os.MkdirAll(destination, 0755)
		if err != nil {
			return utils.FormatError("error creating directory: %v", err)
		}
		isDestinationDir = true
	}

	for _, source := range args[:len(args)-1] {
		finalDestination := destination
		if isDestinationDir {
			// If destination is a directory, append the source filename to the destination directory
			finalDestination = filepath.Join(destination, filepath.Base(source))
		} else if len(args[:len(args)-1]) != 1 {
			return utils.FormatError("mv: target '%s': No such file or directory", destination)
		}

		// Move or rename the file or directory
		err = os.Rename(source, finalDestination)
		if err != nil {
			return utils.FormatError("error moving '%s' to '%s': %v", source, finalDestination, err)
		}
	}

	return nil
}
