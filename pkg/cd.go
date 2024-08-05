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
		usr, err := user.Current()
		if err != nil {
			return fmt.Errorf("\x1b[31mcd: cannot get current user: %s\x1b[0m", err)
		}

		targetDir = usr.HomeDir
	} else {
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

	err := os.Chdir(targetDir)
	if err != nil {
		return fmt.Errorf("cd: %s", err)
	}

	return nil
}
