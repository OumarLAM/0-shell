package pkg

import (
	"fmt"
)

func clear() {
	// ANSI escape code to clear the screen and move the cursor to the home position
	fmt.Print("\033[H\033[2J")
}