package utils

import "fmt"

func FormatError(msg string, args ...interface{}) error {
	return fmt.Errorf("\x1b[31m"+msg+"\x1b[0m", args...)
}