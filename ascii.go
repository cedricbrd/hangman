package hangman

import (
	"bufio"
	"os"
)

func AsciiArt(filePath string) (string, error) {
	art := ""
	data := [][]string{}
	line := 1

	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		if fileScanner.Text() == "" || fileScanner.Text() == " " {
			line++
		} else {
			data = append(data, []string{fileScanner.Text()})
		}
	}

	for i := 0; i < len(data); i++ {
		art += data[i][0] + "\n"
	}

	return art, nil
}

func printHangman(gameState *GameState) string {
	hangman := 10 - gameState.AttemptsLeft
	file, _ := os.Open("position/hangman.txt")
	fileScanner := bufio.NewScanner(file)
	line := 1
	visual := ""
	if hangman == 0 {
		return visual
	} else {
		for fileScanner.Scan() {
			if fileScanner.Text() == "" || fileScanner.Text() == " " {
				line++
			} else if line == hangman {
				visual += "\n"
				visual += fileScanner.Text()
			}
		}
		return visual
	}
}
