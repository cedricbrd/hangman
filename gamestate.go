package hangman

type GameState struct {
	SecretWord   string
	GuessedWord  string
	UsedLetters  string
	AttemptsLeft int
}
