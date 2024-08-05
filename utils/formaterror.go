package utils

import "fmt"

// formatError returns an error message formatted with red color.
func FormatError(msg string, args ...interface{}) error {
	return fmt.Errorf("\x1b[31m"+msg+"\x1b[0m", args...)
}
