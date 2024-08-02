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
	// Get the current user
	usr, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("error getting current user: %s", err)
	}
	username := usr.Username

	// Determine the prompt symbol based on user privileges
	var promptSymbol string
	if username == "root" {
		promptSymbol = "#"
	} else {
		promptSymbol = "$"
	}

	// Get the hostname
	hostname, err := os.Hostname()
	if err != nil {
		return "", fmt.Errorf("error getting hostname: %s", err)
	}

	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting current directory: %s", err)
	}

	// Simplify the directory path to be relative to the user's home directory
	homeDir := usr.HomeDir
	if strings.HasPrefix(cwd, homeDir) {
		cwd = "~" + strings.TrimPrefix(cwd, homeDir)
	}

	// Colorize and format the shell prompt
	// Example colors: username (cyan), hostname (green), directory (blue), prompt symbol (red)
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
	for scanner.Scan() {

		input := scanner.Text()
		args := strings.Fields(input)
		if len(args) > 0 {
			if err := pkg.ExecuteCommand(args); err != nil {
				fmt.Println(err.Error())
			}
		}

		prompt, err := getShellPrompt()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(prompt)
	}
}
