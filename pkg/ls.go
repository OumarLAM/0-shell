package pkg

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/OumarLAM/0-shell/utils"
)

type pathData struct {
	fils            []string
	colone_max_size []int
}

type fileDetails struct {
	mode        os.FileMode
	username    string
	groupName   string
	size        int64
	modTime     string
	path        string
	displayPath string
	fileTypeInd string
	fileCount   int
}

func listDirectory(args []string) error {
	var paths []string
	showAll := false
	longListing := false
	showType := false

	// Determine if the last argument is a path or an option
	// if len(args) > 0 && !strings.HasPrefix(args[len(args)-1], "-") {
	// 	path = args[len(args)-1]  // Set path to the last argument
	// 	args = args[:len(args)-1] // Remove the last argument from args slice
	// } else {
	// 	path = "." // Use current directory if no path is specified
	// }

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
			// // Return an error if there are non-option arguments left
			// return fmt.Errorf("\x1b[31minvalid argument: %s\x1b[0m", arg)
			paths = append(paths, arg)
		}
	}

	if len(paths) == 0 {
		paths = append(paths, ".")
	}

	var allPathsFils = make(map[string]*pathData)

	for _, path := range paths {
		pathinfo := new(pathData)
		// filsPath := []string{}
		// Read the directory contents using os.ReadDir
		files, err := os.ReadDir(path)
		if err != nil {
			return utils.FormatError("ls: %v", err)
		}

		if showAll {
			pathinfo.fils = append(pathinfo.fils, ".")
			dir, err := os.Getwd()
			if err != nil {
				return utils.FormatError("ls: %v", err)
			}
			if dir != "/" {
				pathinfo.fils = append(pathinfo.fils, "..")
			}
		}
		// Display the files based on the options
		for _, entry := range files {
			if !showAll && strings.HasPrefix(entry.Name(), ".") {
				continue // Skip hidden files unless -a is specified
			}
			fullPath := filepath.Join(path, entry.Name()) // Construit le chemin complet
			pathinfo.fils = append(pathinfo.fils, fullPath)
		}

		allPathsFils[path] = pathinfo
	}

	for path, pathinfo := range allPathsFils {
		if len(allPathsFils) != 1 {
			fmt.Println(path + ":")
		}
		updateColumnSizes(pathinfo)
		displayFiles(pathinfo, longListing, showType)
	}
	return nil
}

// fetchFileDetails retrieves all necessary details for a given file
func fetchFileDetails(path string, showType bool) (*fileDetails, error) {
	fullPath := filepath.Join(".", path)
	fi, err := os.Stat(fullPath)
	if err != nil {
		return nil, utils.FormatError("error getting info for %s: %v", fullPath, err)
	}
	stat := fi.Sys().(*syscall.Stat_t)

	usr, err := user.LookupId(fmt.Sprint(stat.Uid))
	if err != nil {
		return nil, utils.FormatError("error looking up user ID %d: %v", stat.Uid, err)
	}
	grp, err := user.LookupGroupId(fmt.Sprint(stat.Gid))
	if err != nil {
		return nil, utils.FormatError("error looking up group ID %d: %v", stat.Gid, err)
	}

	details := &fileDetails{
		mode:        fi.Mode(),
		username:    usr.Username,
		groupName:   grp.Name,
		size:        fi.Size(),
		modTime:     fi.ModTime().Format("Jan 2 15:04"),
		path:        fullPath,
		displayPath: filepath.Base(fullPath),
		fileCount:   1,
	}
	if fi.IsDir() {
		details.displayPath = fmt.Sprintf("\x1b[34m%s\x1b[0m", filepath.Base(fullPath))
		files, err := os.ReadDir(fullPath)
		if err == nil {
			// details.fileCount = len(files) // Compte tous les fichiers et dossiers
			for _, file := range files {
				if file.IsDir() {
					details.fileCount++ // Ajoute un pour chaque sous-dossier pour imiter le comportement de `ls`
				}
			}
			details.fileCount++
		}
	}
	if showType {
		if fi.IsDir() {
			details.fileTypeInd = "/"
		} else if fi.Mode()&os.ModeSymlink != 0 {
			details.fileTypeInd = "@"
		} else if fi.Mode().IsRegular() && (fi.Mode().Perm()&0111) != 0 {
			details.fileTypeInd = "*"
		}
	}

	return details, nil
}

// updateColumnSizes updates the maximum column sizes based on file information
func updateColumnSizes(data *pathData) error {
	if len(data.colone_max_size) < 7 {
		data.colone_max_size = make([]int, 7)
	}

	for _, path := range data.fils {
		details, err := fetchFileDetails(path, false)
		if err != nil {
			return utils.FormatError(err.Error())
		}

		values := []string{
			fmt.Sprintf("%v", details.mode),
			fmt.Sprintf("%d", details.fileCount),
			details.username,
			details.groupName,
			fmt.Sprintf("%d", details.size),
			details.modTime,
			details.path,
		}

		for i, value := range values {
			if len(value) > data.colone_max_size[i] {
				data.colone_max_size[i] = len(value)
			}
		}
	}
	return nil
}

// displayFiles displays file information using the maximum column sizes
func displayFiles(data *pathData, longListing, showType bool) error {
	if longListing {
		total := 0
		for _, path := range data.fils {
			details, err := fetchFileDetails(path, false)
			if err != nil {
				return utils.FormatError(err.Error())
			}
			total += details.fileCount
		}
		fmt.Printf("total %d\n", total)
	}

	for _, path := range data.fils {
		details, err := fetchFileDetails(path, showType)
		if err != nil {
			return utils.FormatError(err.Error())
		}
		// println("****>", details.displayPath,"=", dir+"/")
		// if strings.HasPrefix(details.displayPath, dir+"/") {
		// 	println("===>", details.displayPath)
		// 	name := strings.TrimPrefix(details.displayPath, dir)
		// 	if name != "" {
		// 		details.displayPath = name
		// 	}
		// }

		if longListing {
			fmt.Printf(
				fmt.Sprintf("%%-%dv %%-%dd %%-%ds %%-%ds %%-%dv %%-%ds %%s%%s\n",
					data.colone_max_size[0], data.colone_max_size[1], data.colone_max_size[2],
					data.colone_max_size[3], data.colone_max_size[4], data.colone_max_size[5]),
				details.mode, details.fileCount, details.username, details.groupName, details.size, details.modTime, details.displayPath, details.fileTypeInd)
		} else {
			fmt.Print(details.displayPath + details.fileTypeInd + "  ")
		}
	}
	if !longListing {
		println()
	}
	return nil
}
