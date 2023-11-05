package main

import (
	"github.com/rivo/tview"
)

func main() {
	box := tview.NewFlex().
		SetDirection(tview.FlexRow)
	box.SetBorder(true).
		SetTitle(" Hangman Game ")
	box.SetRect(0, 0, 10, 10)

	box1 := tview.NewBox().
		SetBorder(true).
		SetTitle("2 box")
	box.SetRect(5, 3, 10, 10)
	box2 := tview.NewBox().
		SetBorder(true).
		SetTitle("3 box")
	box.SetRect(0, 0, 10, 10)

	box.
		AddItem(box1, 10, 10, false).
		AddItem(box2, 5, 0, false)

	if err := tview.NewApplication().
		SetRoot(box, true).
		Run(); err != nil {
		panic(err)
	}
}
