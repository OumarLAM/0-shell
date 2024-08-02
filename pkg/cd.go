package pkg

import (
	"fmt"
	"os"
	"os/user"
)

func changeDirectory(args []string) error {
	var targetDir string

	if len(args) < 2 {
		// If no directory is specified, default to the user's home directory
		usr, err := user.Current()
		if err != nil {
			return fmt.Errorf("cd: cannot get current user: %s", err)
		}
		targetDir = usr.HomeDir
	} else {
		// Use the directory specified in the arguments
		targetDir = args[1]
	}

	// Change the directory
	err := os.Chdir(targetDir)
	if err != nil {
		return fmt.Errorf("cd: %s", err)
	}

	return nil
}
