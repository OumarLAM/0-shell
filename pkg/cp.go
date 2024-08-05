package pkg

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func copyFiles(args []string) error {
	destPath := args[len(args)-1]

	for _, sourcePath := range args[:len(args)-1] {
		if err := copy(sourcePath, destPath); err != nil {
			return err
		}
	}

	return nil
}

func copy(source, destination string) error {
	srcInfo, err := os.Stat(source)
	if err != nil {
		return fmt.Errorf("\x1b[31mcp: %v\x1b[0m", err)
	}

	if srcInfo.IsDir() {
		return copyDir(source, destination)
	}

	return copyFile(source, destination)
}

func copyFile(source, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	info, err := os.Stat(destination)
	if err == nil {
		if info.IsDir() {
			destination = filepath.Join(destination, filepath.Base(source))
		}
	} else if !os.IsNotExist(err) {
		if filepath.Ext(destination) == "" {
			err = os.MkdirAll(destination, 0755)
			if err != nil {
				return err
			}
			destination = filepath.Join(destination, filepath.Base(source))
		}
	}

	destFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer destFile.Close()

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
