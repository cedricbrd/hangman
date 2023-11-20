package hangman

import (
	"encoding/json"
	"os"
)

func loadGame(filename string) (GameState, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return GameState{}, err
	}
	var gameState GameState
	err = json.Unmarshal(data, &gameState)
	return gameState, err
}
