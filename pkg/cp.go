package pkg

import (
	"fmt"
	"io"
	"os"
)

func copyFiles(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("\x1b[31mcp: missing file operand\x1b[0m")
	}

	sourcePath := args[0]
	destPath := args[1]

	return copy(sourcePath, destPath)
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
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
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
