package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/OumarLAM/0-shell/pkg"
)

var prompt = "0-Shell> "

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)
	for scanner.Scan() {

		input := scanner.Text()
		args := strings.Fields(input)
		if len(args) > 0 {
			if err := pkg.ExecuteCommand(args); err != nil {
				fmt.Println(err.Error())
			}
		}

		fmt.Print(prompt)
	}
}
