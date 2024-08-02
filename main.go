package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/OumarLAM/0-shell/pkg"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			if err.Error() == "EOF" {
				os.Exit(0)
			}
			continue
		}

		input = strings.TrimSpace(input)

		if input == "exit" {
			os.Exit(0)
		}

		err = pkg.ExecuteCommand(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
