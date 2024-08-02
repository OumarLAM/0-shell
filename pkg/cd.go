package pkg

import "os"

func cd(args []string) error {
	if len(args) == 0 {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		return os.Chdir(homeDir)
	}
	return os.Chdir(args[0])
}
