package pkg

import (
	"fmt"
	"os"
	"strings"
)

func ls(args []string) error {
	var showAll, longFormat, classifyEntries bool
	dir := "."

	for _, arg := range args {
		switch arg {
		case "-a":
			showAll = true
		case "-l":
			longFormat = true
		case "-F":
			classifyEntries = true
		default:
			dir = arg
		}
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		name := entry.Name()
		if !showAll && strings.HasPrefix(name, ".") {
			continue
		}

		if longFormat {
			info, err := entry.Info()
			if err != nil {
				return err
			}
			fmt.Printf("%s %d %s %s\n",
				info.Mode(),
				info.Size(),
				info.ModTime().Format("Jan _2 15:04"),
				info.Name())
		} else {
			if classifyEntries {
				if entry.IsDir() {
					name += "/"
				} else if entry.Type()&os.ModeSymlink != 0 {
					name += "@"
				} else if entry.Type()&os.ModeNamedPipe != 0 {
					name += "|"
				} else if entry.Type()&os.ModeSocket != 0 {
					name += "="
				} else if entry.Type()&os.ModeSetuid != 0 || entry.Type()&os.ModeSetgid != 0 {
					name += "*"
				}
			}
			fmt.Println(name)
		}
	}
	return nil
}
