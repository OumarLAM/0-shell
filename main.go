package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var prompt = "0-Shell> "

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)
	for scanner.Scan() {

		input := scanner.Text()
		args := strings.Fields(input)
		fmt.Println(args)

		fmt.Print(prompt)
	}
}
