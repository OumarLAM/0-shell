package pkg

import (
	"os"
	"time"

	"github.com/OumarLAM/0-shell/utils"
)

// touchFile creates a new file or updates the timestamp of the existing file.
func touchFile(filename string) error {
	// Try to open the file in read-write mode without truncating
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return utils.FormatError(err.Error())
	}
	defer file.Close()

	// Get the current timereturn fmt.Errorf("\x1b[31mrm: %v\x1b[0m", err)
	now := time.Now()

	// Change the access and modification times
	if err := os.Chtimes(filename, now, now); err != nil {
		return utils.FormatError(err.Error())
	}

	return nil
}
