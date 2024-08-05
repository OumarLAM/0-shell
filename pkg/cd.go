package pkg

import (
	"os"
	"os/user"
	"strings"

	"github.com/OumarLAM/0-shell/utils"
)

func changeDirectory(args []string) error {
	var targetDir string

	if len(args) < 1 {
		usr, err := user.Current()
		if err != nil {
			return utils.FormatError("cd: cannot get current user: %s", err)
		}
		targetDir = usr.HomeDir
	} else {
		if strings.Contains(args[0], "~") {
			usr, err := user.Current()
			if err != nil {
				return utils.FormatError("cd: cannot get current user: %s", err)
			}
			args[0] = strings.ReplaceAll(args[0], "~", usr.HomeDir)
		}
		targetDir = args[0]
	}

	err := os.Chdir(targetDir)
	if err != nil {
		return utils.FormatError("cd: %s", err)
	}

	return nil
}