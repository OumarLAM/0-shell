package pkg

import (
	"io"
	"os"
	"path/filepath"

	"github.com/OumarLAM/0-shell/utils"
)

func copyFiles(args []string) error {

	// sourcePath := args[0]
	destPath := args[len(args)-1]
	destinationInfo, err := os.Stat(destPath)
	// Check if the destination exists and is a directory
	isDestinationDir := false
	if err == nil {
		if destinationInfo.IsDir() {
			isDestinationDir = true
		}
	} else if !os.IsNotExist(err) {
		return utils.FormatError("error accessing destination: %v", err)
	} else if filepath.Ext(destPath) == "" {
		// If the destination does not exist and has no extension, treat as a directory
		err = os.MkdirAll(destPath, 0755)
		if err != nil {
			return utils.FormatError("error creating directory: %v", err)
		}
		isDestinationDir = true
	}

	for _, sourcePath := range args[:len(args)-1] {
		if len(args[:len(args)-1]) != 1 && !isDestinationDir {
			return utils.FormatError("cp: target '%s': No such file or directory", destPath)
		}
		if err := copy(sourcePath, destPath); err != nil {
			return err
		}
	}

	return nil
}

func copy(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return utils.FormatError("cp: %v", err)
	}

	if srcInfo.IsDir() {
		return copyDir(src, dst)
	}

	return copyFile(src, dst)
}

func copyFile(src, dst string) error {
	// Open the source file
	sourceFile, err := os.Open(src)
	if err != nil {
		return utils.FormatError(err.Error())
	}
	defer sourceFile.Close()

	// Check if the destination exists and determine the nature
	info, err := os.Stat(dst)
	if err == nil {
		if info.IsDir() {
			// If dst is a directory, append the src filename to the dst directory
			dst = filepath.Join(dst, filepath.Base(src))
		}
	} else if !os.IsNotExist(err) {
		return utils.FormatError("error accessing destination: %v", err)
	} else if filepath.Ext(dst) == "" {
		// No extension, treat as a directory
		err = os.MkdirAll(dst, 0755)
		if err != nil {
			return utils.FormatError(err.Error())
		}
		// Set the destination to a new file inside the newly created directory
		dst = filepath.Join(dst, filepath.Base(src))
	}

	// Create the destination file
	destFile, err := os.Create(dst)
	if err != nil {
		return utils.FormatError(err.Error())
	}
	defer destFile.Close()

	// Perform the file copy
	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return utils.FormatError(err.Error())
	}
	return nil
}

func copyDir(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return utils.FormatError(err.Error())
	}

	err = os.MkdirAll(dst, srcInfo.Mode())
	if err != nil {
		return utils.FormatError(err.Error())
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return utils.FormatError(err.Error())
	}

	for _, entry := range entries {
		srcPath := src + "/" + entry.Name()
		dstPath := dst + "/" + entry.Name()

		err = copy(srcPath, dstPath)
		if err != nil {
			return utils.FormatError(err.Error())
		}
	}

	return nil
}
