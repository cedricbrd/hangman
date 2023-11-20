package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

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

	if len(mots) == 0 {
		return "", fmt.Errorf("le fichier est vide")
	}

	random := rand.New(rand.NewSource(time.Now().Unix()))
	motChoisi := mots[random.Intn(len(mots))]
	return motChoisi, nil
}
