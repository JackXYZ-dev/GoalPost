package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	textView := tview.NewTextView().
		SetText("Welcome to GoalPost!").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	if err := app.SetRoot(textView, true).Run(); err != nil {
		panic(err)
	}
}