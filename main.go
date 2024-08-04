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

	var multilineInput string
	var inQuotes bool
	var quoteChar rune

	for scanner.Scan() {
		input := scanner.Text()
		// Vérifier chaque caractère pour gérer les guillemets
		for i, char := range input {
			if (char == '"' || char == '\'') && (i == 0 || input[i-1] != '\\') {
				if inQuotes && char == quoteChar {
					inQuotes = false // Fermeture du guillemet trouvé
				} else if !inQuotes {
					inQuotes = true
					quoteChar = char
				}
			}
		}

		if inQuotes {
			multilineInput += input + "\n" // Ajouter un saut de ligne à chaque ligne ajoutée
			fmt.Print("> ")                // Prompt pour plus d'entrée
			continue
		} else {
			multilineInput += input // Ajouter la dernière ligne sans saut de ligne supplémentaire
		}

		// Traitement de la commande complète
		if !inQuotes { // Supprimer les guillemets uniquement si tous les guillemets sont fermés
			multilineInput = strings.Replace(multilineInput, string(quoteChar), "", -1)
		}

		args := strings.Fields(multilineInput)
		if len(args) > 0 {
			if err := pkg.ExecuteCommand(args, multilineInput); err != nil {
				fmt.Println(err.Error())
			}
		}

		multilineInput = "" // Réinitialiser pour la prochaine commande
		prompt, err = getShellPrompt()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(prompt)
	}
}
