package pkg

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/OumarLAM/0-shell/utils"
)

func ParseAndApplyChmod(mode string, files []string) error {
	// var err error
	if strings.ContainsAny(mode, "+-=") {
		// Handle symbolic mode
		return utils.FormatError("chmod: invalid mode: %s", mode)
		// for _, file := range files {
		// 	fileInfo, err := os.Stat(file)
		// 	if err != nil {
		// 		return fmt.Errorf("chmod: cannot access '%s': %v", file, err)
		// 	}
		// 	currentMode := fileInfo.Mode()
		// 	newMode, err := applySymbolicMode(mode, currentMode)
		// 	println(newMode, err)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	err = os.Chmod(file, newMode)
		// 	if err != nil {
		// 		return fmt.Errorf("chmod: cannot change permissions of '%s': %v", file, err)
		// 	}
		// }
	} else {
		// Handle numeric mode
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

func applySymbolicMode(mode string, currentMode os.FileMode) (os.FileMode, error) {
	// Split the mode string on commas to handle multiple changes
	modeChanges := strings.Split(mode, ",")
	for _, change := range modeChanges {
		if len(change) < 2 { // Must at least have an operation and a permission
			return currentMode, fmt.Errorf("invalid mode specification: %s", change)
		}

		// Determine the operation (+, -, =)
		operation := change[0]
		if operation != '+' && operation != '-' && operation != '=' {
			return currentMode, fmt.Errorf("invalid operation: %c", operation)
		}

		// Determine the target (u, g, o, a)
		targets := change[1 : len(change)-1]
		permission := change[len(change)-1]

		// Map permission character to FileMode bit
		var permBits os.FileMode
		switch permission {
		case 'r':
			permBits = 0444
		case 'w':
			permBits = 0222
		case 'x':
			permBits = 0111
		default:
			return currentMode, fmt.Errorf("invalid permission: %c", permission)
		}

		// Apply the change based on the target
		for _, target := range targets {
			switch target {
			case 'u':
				if operation == '+' {
					currentMode |= (permBits & 0700)
				} else if operation == '-' {
					currentMode &^= (permBits & 0700)
				} else if operation == '=' {
					currentMode &^= 0700 // Clear user bits
					currentMode |= (permBits & 0700)
				}
			case 'g':
				if operation == '+' {
					currentMode |= (permBits & 0070)
				} else if operation == '-' {
					currentMode &^= (permBits & 0070)
				} else if operation == '=' {
					currentMode &^= 0070 // Clear group bits
					currentMode |= (permBits & 0070)
				}
			case 'o':
				if operation == '+' {
					currentMode |= (permBits & 0007)
				} else if operation == '-' {
					currentMode &^= (permBits & 0007)
				} else if operation == '=' {
					currentMode &^= 0007 // Clear other bits
					currentMode |= (permBits & 0007)
				}
			case 'a':
				if operation == '+' {
					currentMode |= permBits
				} else if operation == '-' {
					currentMode &^= permBits
				} else if operation == '=' {
					currentMode &^= 0777 // Clear all bits
					currentMode |= permBits
				}
			default:
				return currentMode, fmt.Errorf("invalid target: %c", target)
			}
		}
	}

	return currentMode, nil
}
