package pkg

import (
	"fmt"
	"os"
	"path/filepath"
)

func moveFiles(args []string) error {

	// source := args[0]
	destination := args[len(args)-1]

	for _, source := range args[:len(args)-1] {
		info, err := os.Stat(destination)
		if err == nil {
			if info.IsDir() {
				// If dst is a directory, append the src filename to the dst directory
				destination = filepath.Join(destination, filepath.Base(source))
			}
		} else if !os.IsNotExist(err) {
			if filepath.Ext(destination) == "" {
				// No extension, treat as a directory
				err = os.MkdirAll(destination, 0755)
				if err != nil {
					return err
				}
				// Set the destination to a new file inside the newly created directory
				destination = filepath.Join(destination, filepath.Base(destination))
			}
		}
		// Move or rename the file or directory
		err = os.Rename(source, destination)
		if err != nil {
			return fmt.Errorf("\x1b[31mmv: %v\x1b[0m", err)
		}
	}

	return nil
}
