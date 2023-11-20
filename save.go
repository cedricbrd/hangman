package hangman

import (
	"encoding/json"
	"os"
)

func saveGame(filename string, gameState GameState) error {
	data, err := json.Marshal(gameState)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	return err
}
