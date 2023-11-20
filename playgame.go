package hangman

import (
	"fmt"
	"os"
	"strings"
)

func PlayGame() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: " + os.Args[0] + "")
		return
	}

	wordFile := os.Args[1]
	if !strings.HasSuffix(wordFile, ".txt") {
		fmt.Println("Format de fichier invalide, veuillez choisir un fichier au format .txt")
		return
	}

	var choix int
	fmt.Println("1. Nouvelle Partie")
	fmt.Println("2. Charger une Partie")
	fmt.Print("|---------------|\n")

	_, err := fmt.Scanln(&choix)
	if err != nil || (choix != 1 && choix != 2) {
		fmt.Println("Choix invalide")
		return
	}

	gameState := &GameState{}

	if choix == 2 {
		var saveFile string
		fmt.Print("Entrez le nom du fichier de sauvegarde : ")
		_, err := fmt.Scanln(&saveFile)
		if err != nil {
			fmt.Println("Erreur de saisie")
			return
		}

		savedGame, err := loadGame(saveFile)
		if err != nil {
			fmt.Println("Erreur lors du chargement de la partie :", err)
			return
		}

		gameState = &GameState{
			SecretWord:   savedGame.SecretWord,
			GuessedWord:  savedGame.GuessedWord,
			UsedLetters:  savedGame.UsedLetters,
			AttemptsLeft: savedGame.AttemptsLeft,
		}
	} else {
		fichierMots := wordFile

		secretWord, err := randomWord(fichierMots)
		if err != nil {
			fmt.Println("Erreur lors de la lecture du fichier de mots : ", err)
			return
		}

		attempts := 10
		guessedWord := strings.Repeat("_", len(secretWord))
		UsedLetters := ""

		fmt.Println("Début de la partie")
		gameState = &GameState{
			SecretWord:   secretWord,
			GuessedWord:  guessedWord,
			UsedLetters:  UsedLetters,
			AttemptsLeft: attempts,
		}
	}

	for gameState.AttemptsLeft > 0 {
		clearTerminal()
		fmt.Println("Mot à deviner : ")
		fmt.Println(printHangman(gameState))
		fmt.Println(gameState.GuessedWord)
		fmt.Println("Lettres déjà utilisées : ", gameState.UsedLetters)
		fmt.Println("Tentatives restantes : ", gameState.AttemptsLeft)

		var guess string
		fmt.Print("Devinez une lettre ou le mot complet : ")
		_, err := fmt.Scanln(&guess)
		if err != nil {
			fmt.Println("Erreur de saisie")
			continue
		}

		if guess == "quit" || guess == "Quit" {
			clearTerminal()
			os.Exit(0)
		}

		if guess == "save" || guess == "Save" {
			err = saveGame("save.json", GameState{
				SecretWord:   gameState.SecretWord,
				GuessedWord:  gameState.GuessedWord,
				UsedLetters:  gameState.UsedLetters,
				AttemptsLeft: gameState.AttemptsLeft,
			})
			if err != nil {
				fmt.Println("Erreur lors de la sauvegarde de la partie :", err)
				return
			}
			fmt.Println("Partie sauvegardée avec succès")
			break
		}

		if len(guess) > 1 {
			if guess == gameState.SecretWord {
				gameState.GuessedWord = gameState.SecretWord
				fmt.Printf("Vous avez deviné le mot ! Le mot était %s.\n", gameState.SecretWord)

				err := os.Remove("save.json")
				if err != nil {
					fmt.Println("", err)
				}

				break
			} else {
				fmt.Println("Le mot est incorrect!")
				gameState.AttemptsLeft--
			}
		} else if len(guess) == 1 && isLetter(guess) {
			usedLetters := guess[0]
			fmt.Println("La lettre est correcte ou incorrecte")
			gameState.UsedLetters += string(usedLetters)

			newGuessedWord := ""
			for _, letter := range gameState.SecretWord {
				if strings.ContainsRune(gameState.UsedLetters, letter) {
					newGuessedWord += string(letter)
				} else {
					newGuessedWord += "_"
				}
			}
			gameState.GuessedWord = newGuessedWord

			if gameState.GuessedWord == gameState.SecretWord {
				fmt.Printf("Vous avez deviné le mot ! Le mot était %s.\n", gameState.SecretWord)

				err := os.Remove("save.json")
				if err != nil {
					fmt.Println("Erreur lors de la suppression du fichier du sauvegarde :", err)
				}

				break
			} else {
				fmt.Println("La lettre choisie est incorrecte")
				gameState.AttemptsLeft--
				printHangman(gameState)
			}

		} else {
			fmt.Println("Veuillez entrer une seule lettre valide")
		}
	}

	if gameState.AttemptsLeft == 0 {
		clearTerminal()
		printHangman(gameState)
		fmt.Printf("Vous avez perdu ! Le mot était %s.\n", gameState.SecretWord)

		err := os.Remove("save.json")
		if err != nil {
			fmt.Println("")
		}
	}
}

func clearTerminal() {
	fmt.Print("\033[H\033[2J")
}
