package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/OumarLAM/0-shell/pkg"
)

func getShellPrompt() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("error getting current user: %s", err)
	}
	username := usr.Username

	var promptSymbol string
	if username == "root" {
		promptSymbol = "#"
	} else {
		promptSymbol = "$"
	}

	hostname, err := os.Hostname()
	if err != nil {
		return "", fmt.Errorf("error getting hostname: %s", err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting current directory: %s", err)
	}

	homeDir := usr.HomeDir
	if strings.HasPrefix(cwd, homeDir) {
		cwd = "~" + strings.TrimPrefix(cwd, homeDir)
	}

	prompt := fmt.Sprintf("\x1b[36m%s\x1b[0m@\x1b[32m%s\x1b[0m:\x1b[34m%s\x1b[0m\x1b[31m%s \x1b[0m", username, hostname, cwd, promptSymbol)

	return prompt, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	prompt, err := getShellPrompt()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(prompt)

	var multilineInput string
	var inQuotes bool
	var quoteChar rune

	for scanner.Scan() {
		input := scanner.Text()
		for i, char := range input {
			if (char == '"' || char == '\'') && (i == 0 || input[i-1] != '\\') {
				if inQuotes && char == quoteChar {
					inQuotes = false
				} else if !inQuotes {
					inQuotes = true
					quoteChar = char
				}
			}
		}

		if inQuotes {
			multilineInput += input + "\n"
			fmt.Print("> ")
			continue
		} else {
			multilineInput += input
		}

		if !inQuotes {
			multilineInput = strings.Replace(multilineInput, string(quoteChar), "", -1)
		}

		args := strings.Fields(multilineInput)
		if len(args) > 0 {
			if err := pkg.ExecuteCommand(args, multilineInput); err != nil {
				fmt.Println(err.Error())
			}
		}

		multilineInput = ""
		prompt, err = getShellPrompt()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(prompt)
	}
}