package pkg

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

func changeDirectory(args []string) error {
	var targetDir string

	if len(args) < 1 {
		// If no directory is specified, default to the user's home directory
		usr, err := user.Current()
		if err != nil {
			return fmt.Errorf("cd: cannot get current user: %s", err)
		}
		targetDir = usr.HomeDir
	} else {
		// Use the directory specified in the arguments
		if strings.Contains(args[0], "~") {
			usr, err := user.Current()
			if err != nil {
				return fmt.Errorf("cd: cannot get current user: %s", err)
			}
			args[0] = strings.ReplaceAll(args[0], "~", usr.HomeDir)
		}
		targetDir = args[0]
	}

	fmt.Println(targetDir)
	// Change the directory
	err := os.Chdir(targetDir)
	if err != nil {
		return fmt.Errorf("cd: %s", err)
	}

	return nil
}
