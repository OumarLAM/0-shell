package pkg

import (
	"fmt"
	"os"
	"strings"
)

func listDirectory(args []string) error {
	var path string
	showAll := false
	longListing := false
	showType := false

	// Determine if the last argument is a path or an option
	if len(args) > 0 && !strings.HasPrefix(args[len(args)-1], "-") {
		path = args[len(args)-1]  // Set path to the last argument
		args = args[:len(args)-1] // Remove the last argument from args slice
	} else {
		path = "." // Use current directory if no path is specified
	}

	// Parse the options
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			if strings.Contains(arg, "a") {
				showAll = true
			}
			if strings.Contains(arg, "l") {
				longListing = true
			}
			if strings.Contains(arg, "F") {
				showType = true
			}
		} else {
			// Return an error if there are non-option arguments left
			return fmt.Errorf("\x1b[31minvalid argument: %s\x1b[0m", arg)
		}
	}

	// Read the directory contents using os.ReadDir
	files, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("\x1b[31mls: %v\x1b[0m", err)
	}

	if showAll {
		fmt.Println(".")
		fmt.Println("..")
	}
	// Display the files based on the options
	for _, entry := range files {
		if !showAll && strings.HasPrefix(entry.Name(), ".") {
			continue // Skip hidden files unless -a is specified
		}

		fileInfo := entry.Name()
		info, err := entry.Info()
		if err != nil {
			return fmt.Errorf("\x1b[31merror retrieving file info: %v\x1b[0m", err)
		}

		if showType {
			if entry.IsDir() {
				fileInfo += "/"
			} else if info.Mode()&os.ModeSymlink != 0 {
				fileInfo += "@"
			} else if info.Mode().IsRegular() && info.Mode().Perm()&0111 != 0 {
				fileInfo += "*"
			}
		}

		if longListing {
			fmt.Printf("%v %v %v %v\n", info.Mode(), info.Size(), info.ModTime().Format("Jan 02 15:04"), fileInfo)
		} else {
			fmt.Println(fileInfo)
		}
	}

	return nil
}
