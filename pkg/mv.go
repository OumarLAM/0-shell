package pkg

import (
	"fmt"
	"os"
	"path/filepath"
)

func moveFiles(args []string) error {
	destination := args[len(args)-1]

	for _, source := range args[:len(args)-1] {
		info, err := os.Stat(destination)
		if err == nil {
			if info.IsDir() {
				destination = filepath.Join(destination, filepath.Base(source))
			}
		} else if !os.IsNotExist(err) {
			if filepath.Ext(destination) == "" {
				err = os.MkdirAll(destination, 0755)
				if err != nil {
					return err
				}
				destination = filepath.Join(destination, filepath.Base(destination))
			}
		}
		
		err = os.Rename(source, destination)
		if err != nil {
			return fmt.Errorf("\x1b[31mmv: %v\x1b[0m", err)
		}
	}

	return nil
}
