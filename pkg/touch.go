package pkg

import (
	"os"
	"time"

	"github.com/OumarLAM/0-shell/utils"
)

func touchFile(filename string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return utils.FormatError(err.Error())
	}
	defer file.Close()

	now := time.Now()

	if err := os.Chtimes(filename, now, now); err != nil {
		return utils.FormatError(err.Error())
	}

	return nil
}