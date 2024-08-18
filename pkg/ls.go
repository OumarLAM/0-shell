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
	blockSize   int64
}

func listDirectory(args []string) error {
	var paths []string
	showAll := false
	longListing := false
	showType := false

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
			paths = append(paths, arg)
		}
	}

	if len(paths) == 0 {
		paths = append(paths, ".")
	}

	var allPathsFils = make(map[string]*pathData)

	for _, path := range paths {
		pathinfo := new(pathData)
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

		for _, entry := range files {
			if !showAll && strings.HasPrefix(entry.Name(), ".") {
				continue
			}
			fullPath := filepath.Join(path, entry.Name())
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

func fetchFileDetails(path string, showType bool) (*fileDetails, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, utils.FormatError("error getting info for %s: %v", path, err)
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
		path:        path,
		fileCount:   1,
		displayPath: path,
		blockSize:   int64(stat.Blocks),
	}

	if fi.IsDir() {
		details.displayPath = fmt.Sprintf("\x1b[34m%s\x1b[0m", path)
		files, err := os.ReadDir(path)
		if err == nil {
			for _, file := range files {
				if file.IsDir() {
					details.fileCount++
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

func displayFiles(data *pathData, longListing, showType bool) error {
	if longListing {
		totalBlocks := int64(0)
		for _, path := range data.fils {
			details, err := fetchFileDetails(path, false)
			if err != nil {
				return utils.FormatError(err.Error())
			}
			totalBlocks += details.blockSize
		}
		fmt.Printf("total %d\n", totalBlocks/2)
	}

	for _, path := range data.fils {
		details, err := fetchFileDetails(path, showType)
		if err != nil {
			return utils.FormatError(err.Error())
		}

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
