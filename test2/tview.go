package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	// Créez un Box avec une bordure et un titre
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow)

	flex.SetBorder(true).
		SetTitle(" Hangman Game ")
	flex.SetRect(0, 0, 10, 10)

	box2 := tview.NewBox().SetBorder(true).SetTitle("box2")

	// Définissez la position et la taille du Box
	flex.AddItem(box2, 0, 0, false)
	box2.SetRect(0, 0, 0, 0)

	flex.
		AddItem(box2, 10, 10, false)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
