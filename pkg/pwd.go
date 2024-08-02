package pkg

import (
	"fmt"
	"os"
)

func Pwd() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(dir)
	return nil
}
