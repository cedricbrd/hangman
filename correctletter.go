package hangman

import "strings"

func CorrectLetter(correctLetters byte, word string) bool {
	correctLetters = byte(strings.ToLower(string(correctLetters))[0])

	for _, lettreMot := range word {
		lettreMot = rune(strings.ToLower(string(lettreMot))[0])
		if lettreMot == rune(correctLetters) {
			return true
		}
	}
	return false
}
