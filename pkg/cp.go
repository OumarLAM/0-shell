package pkg

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func copyFiles(args []string) error {

	// sourcePath := args[0]
	destPath := args[len(args)-1]

	for _, sourcePath := range args[:len(args)-1] {
		if err := copy(sourcePath, destPath); err != nil {
			return err
		}
	}

	return nil
}

func copy(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("\x1b[31mcp: %v\x1b[0m", err)
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
		return err
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
		if filepath.Ext(dst) == "" {
			// No extension, treat as a directory
			err = os.MkdirAll(dst, 0755)
			if err != nil {
				return err
			}
			// Set the destination to a new file inside the newly created directory
			dst = filepath.Join(dst, filepath.Base(src))
		}
	}

	// Create the destination file
	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Perform the file copy
	_, err = io.Copy(destFile, sourceFile)
	return err
}

func copyDir(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dst, srcInfo.Mode())
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := src + "/" + entry.Name()
		dstPath := dst + "/" + entry.Name()

		err = copy(srcPath, dstPath)
		if err != nil {
			return err
		}
	}

	return nil
}
