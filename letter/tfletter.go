package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/01-edu/z01"
)

type GameState struct {
	SecretWord     string
	GuessedWord    string
	guessedLetters string
	attemptsLeft   int
}

func correctLetter(lettreTrouvee byte, motADeviner string) bool {
	lettreTrouvee = byte(strings.ToLower(string(lettreTrouvee))[0])

	for _, lettreMot := range motADeviner {
		lettreMot = rune(strings.ToLower(string(lettreMot))[0])
		if lettreMot == rune(lettreTrouvee) {
			return true
		}
	}
	return false
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

	rand.Seed(time.Now().Unix())
	motChoisi := mots[rand.Intn(len(mots))]
	return motChoisi, nil
}

func main() {

	files, err := listFilesInDirectory("dictionnary", ".txt")
	if err != nil {
		fmt.Println("Erreur lors de la lecture des fichiers: ", err)
		return
	}

	fmt.Println("Veuillez choisir le fichier de mot de votre choix :")
	for i, file := range files {
		fmt.Printf("%d. %s\n", i+1, file)
	}

	var choix int
	fmt.Print("|---------------|")
	z01.PrintRune('\n')
	_, err = fmt.Scanln(&choix)
	if err != nil || choix < 1 || choix > len(files) {
		fmt.Println("Choix non valide.")
		return
	}

	fichierMots := filepath.Join("dictionnary", files[choix-1])

	secretWord, err := randomWord(fichierMots)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier de mots: ", err)
		return
	}

	attempts := 6
	guessedWord := strings.Repeat("_", len(secretWord))
	guessedLetters := ""

	fmt.Println("C'est parti!")

	for attempts > 0 {
		fmt.Printf("Mot à deviner: %s\n", guessedWord)
		fmt.Printf("Lettres déjà devinées: %s\n", guessedLetters)
		fmt.Printf("Tentatives restantes: %d\n", attempts)

		var letter string
		fmt.Print("Devinez une lettre: ")
		_, err := fmt.Scanln(&letter)
		if err != nil {
			fmt.Println("Erreur de saisie: ", err)
			continue
		}

		if len(letter) != 1 || !isLetter(letter) {
			fmt.Println("Veuillez entrer une seule lettre valide.")
			continue
		}

		guessedLetter := letter[0]

		if correctLetter(byte(guessedLetter), secretWord) {
			fmt.Println("La lettre est correcte!")
			guessedLetters += string(guessedLetter)

			newGuessedWord := ""
			for _, letter := range secretWord {
				if strings.ContainsRune(guessedLetters, letter) {
					newGuessedWord += string(letter)
				} else {
					newGuessedWord += "_"
				}
			}
			guessedWord = newGuessedWord

			if guessedWord == secretWord {
				fmt.Printf("Vous avez deviné le mot! Le mot était %s.\n", secretWord)
				break
			}
		} else {
			fmt.Println("La lettre est incorrecte!")
			attempts--
		}
	}

	if attempts == 0 {
		fmt.Printf("Vous avez perdu! Le mot était %s.\n", secretWord)
	}
}

func isLetter(s string) bool {
	return len(s) == 1 && s >= "a" && s <= "z" || s >= "A" && s <= "Z"
}

func listFilesInDirectory(directory, extension string) ([]string, error) {
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

func saveGame(gameState GameState) error {
	// Ouvrir le fichier pour écriture
	file, err := os.Create("saved_game.json")
	if err != nil {
		return err
	}
	defer file.Close()

	// Encoder la structure GameState en JSON
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(gameState); err != nil {
		return err
	}

	fmt.Println("Partie sauvegardée avec succès!")
	return nil
}
