package pkg

import (
	"fmt"
	"strings"
)

func echo(args []string) error {
	fmt.Println(strings.Join(args, " "))
	return nil
}