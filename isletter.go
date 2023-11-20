package hangman

func isLetter(s string) bool {
	return len(s) == 1 && s >= "a" && s <= "z" || s >= "A" && s <= "Z"
}
