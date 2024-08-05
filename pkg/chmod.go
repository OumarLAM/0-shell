package pkg

import (
	"os"
	"strconv"
	"strings"

	"github.com/OumarLAM/0-shell/utils"
)

func ParseAndApplyChmod(mode string, files []string) error {
	if strings.ContainsAny(mode, "+-=") {
		return utils.FormatError("chmod: invalid mode: %s", mode)
	} else {
		parsedMode, err := strconv.ParseUint(mode, 8, 32)
		if err != nil {
			return utils.FormatError("chmod: invalid mode: %s", mode)
		}
		for _, file := range files {
			err = os.Chmod(file, os.FileMode(parsedMode))
			if err != nil {
				return utils.FormatError("chmod: cannot change permissions of '%s': %v", file, err)
			}
		}
	}
	return nil
}