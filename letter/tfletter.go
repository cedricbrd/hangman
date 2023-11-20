package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/01-edu/z01"
)

type GameState struct {
	SecretWord   string
	GuessedWord  string
	UsedLetters  string
	AttemptsLeft int
}

type Save struct {
	SecretWord   string
	GuessedWord  string
	UsedLetters  string
	AttemptsLeft int
}

func HangmanAscii(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var hangmanAscii []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hangmanAscii = append(hangmanAscii, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return hangmanAscii, nil
}

func randomWord(fichierMots string) (string, error) {
	file, err := os.Open(fichierMots)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var mots []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mots = append(mots, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}
	random := rand.New(rand.NewSource(time.Now().Unix()))
	motChoisi := mots[random.Intn(len(mots))]
	return motChoisi, nil
}

func searchFile(directory, extension string) ([]string, error) {
	var files []string
	dir, err := os.Open(directory)
	if err != nil {
		return nil, err
	}
	defer dir.Close()
	fileInfos, err := dir.Readdir(0)
	if err != nil {
		return nil, err
	}
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() && strings.HasSuffix(fileInfo.Name(), extension) {
			files = append(files, fileInfo.Name())
		}
	}
	return files, nil
}

func saveGame(filename string, gameState Save) error {
	data, err := json.Marshal(gameState)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	return err
}

func loadGame(filename string) (Save, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return Save{}, err
	}
	var gameState Save
	err = json.Unmarshal(data, &gameState)
	return gameState, err
}

func main() {
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

	hangmanAscii, err := HangmanAscii("position/hangman.txt")
	if err != nil {
		fmt.Println("Erreur lors du lancement du fchier Ascii :", err)
		return
	}

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

		gameState := &GameState{
			SecretWord:   savedGame.SecretWord,
			GuessedWord:  savedGame.GuessedWord,
			UsedLetters:  savedGame.UsedLetters,
			AttemptsLeft: savedGame.AttemptsLeft,
		}

		playGame(gameState, hangmanAscii)
	} else {
		fichierMots := wordFile

		secretWord, err := randomWord(fichierMots)
		if err != nil {
			fmt.Println("Erreur lors de la lecture du fichier de mots : ", err)
			return
		}

		attempts := 10
		guessedWord := strings.Repeat("(__________)", len(secretWord))
		UsedLetters := ""

		fmt.Println("Début de la partie")
		gameState := &GameState{
			SecretWord:   secretWord,
			GuessedWord:  guessedWord,
			UsedLetters:  UsedLetters,
			AttemptsLeft: attempts,
		}

		playGame(gameState, hangmanAscii)
	}
}

func correctLetter(correctLetters byte, word string) bool {
	correctLetters = byte(strings.ToLower(string(correctLetters))[0])

	for _, lettreMot := range word {
		lettreMot = rune(strings.ToLower(string(lettreMot))[0])
		if lettreMot == rune(correctLetters) {
			return true
		}
	}
	return false
}

func isLetter(s string) bool {
	return len(s) == 1 && s >= "a" && s <= "z" || s >= "A" && s <= "Z"
}

func playGame(gameState *GameState, hangmanAscii []string) {
	for gameState.AttemptsLeft > 0 {
		clearTerminal()
		printHangman(gameState, hangmanAscii)

		fmt.Println("Mot à deviner : ")
		z01.PrintRune('\n')
		z01.PrintRune('\n')
		z01.PrintRune('\n')
		z01.PrintRune('\n')
		z01.PrintRune('\n')
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

		usedLetters := guess[0]
		gameState.UsedLetters += string(usedLetters)

		if guess == "quit" {
			break
		}

		if guess == "save" {
			err = saveGame("save.json", Save{
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
					fmt.Println("Erreur lors de la suppression du fichier de sauvegarde :", err)
				}

				break
			} else {
				fmt.Println("Le mot est incorrect!")
				gameState.AttemptsLeft--
			}
		} else if len(guess) == 1 && isLetter(guess) {
			usedLetters := guess[0]
			if correctLetter(byte(usedLetters), gameState.SecretWord) {

				newGuessedWord := ""
				for _, letter := range gameState.SecretWord {
					if strings.ContainsRune(gameState.UsedLetters, letter) {
						newGuessedWord += string(letter)
					} else {
						newGuessedWord += "(__________)"
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
				}
			} else {
				fmt.Println("La lettre choisie est incorrecte")
				gameState.AttemptsLeft--
				printHangman(gameState, hangmanAscii)
			}
		} else {
			fmt.Println("Veuillez entrer une seule lettre valide")
		}
	}

	if gameState.AttemptsLeft == 0 {
		clearTerminal()
		printHangman(gameState, hangmanAscii)
		fmt.Printf("Vous avez perdu ! Le mot était %s.\n", gameState.SecretWord)

		err := os.Remove("save.json")
		if err != nil {
			fmt.Println("")
		}
	}
}

func AsciiArt(str string) string {
	art := ""
	data := [][]string{}
	line := 1

	for e := range str {
		data = append(data, []string{})
		line = 1
		file, _ := os.Open("ASCII/standard.txt")
		fileScanner := bufio.NewScanner(file)
		for fileScanner.Scan() {
			if fileScanner.Text() == "" || fileScanner.Text() == " " {
				line++
			} else if line == int(rune(str[e]))-30 {
				data[e] = append(data[e], fileScanner.Text())
			}
		}
		file.Close()
	}
	for i := 0; i < 8; i++ {
		for e := range data {
			art += data[e][i]
		}
		art += "\n"
	}
	return art
}

func clearTerminal() {
	fmt.Print("\033[H\033[2J")
}

func printHangman(gameState *GameState, hangmanAscii []string) {
	index := 10 - gameState.AttemptsLeft
	if index >= 0 && index < len(hangmanAscii) {
		fmt.Println(hangmanAscii[index])
	} else {
		fmt.Println("Invalid index for Hangman ASCII art")
	}
}
