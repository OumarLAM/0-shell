package pkg

import (
	"fmt"
	"os"

	"github.com/OumarLAM/0-shell/utils"
)

func concatenateFiles(args []string) error {
	if len(args) == 0 {
		return utils.FormatError("no input files")
	}

	for _, filename := range args {
		data, err := os.ReadFile(filename)
		if err != nil {
			return utils.FormatError("cat: %v", err)
		}
		fmt.Print(string(data))
	}
	println()
	return nil
}