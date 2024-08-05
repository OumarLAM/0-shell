package pkg

import (
	"io"
	"os"
	"path/filepath"

	"github.com/OumarLAM/0-shell/utils"
)

func copyFiles(args []string) error {

	destPath := args[len(args)-1]
	destinationInfo, err := os.Stat(destPath)
	isDestinationDir := false
	if err == nil {
		if destinationInfo.IsDir() {
			isDestinationDir = true
		}
	} else if !os.IsNotExist(err) {
		return utils.FormatError("error accessing destination: %v", err)
	} else if filepath.Ext(destPath) == "" {
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
	sourceFile, err := os.Open(src)
	if err != nil {
		return utils.FormatError(err.Error())
	}
	defer sourceFile.Close()

	info, err := os.Stat(dst)
	if err == nil {
		if info.IsDir() {
			dst = filepath.Join(dst, filepath.Base(src))
		}
	} else if !os.IsNotExist(err) {
		return utils.FormatError("error accessing destination: %v", err)
	} else if filepath.Ext(dst) == "" {
		err = os.MkdirAll(dst, 0755)
		if err != nil {
			return utils.FormatError(err.Error())
		}
		dst = filepath.Join(dst, filepath.Base(src))
	}

	destFile, err := os.Create(dst)
	if err != nil {
		return utils.FormatError(err.Error())
	}
	defer destFile.Close()

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
