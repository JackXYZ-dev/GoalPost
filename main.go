package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	menu := tview.NewList().
		AddItem("Fixtures", "View upcoming matches", 'f', nil).
		AddItem("Results", "View recent results", 'r', nil).
		AddItem("Tables", "View league tables", 't', nil).
		AddItem("News", "View latest news", 'n', nil).
		AddItem("Quit", "Exit the program", 'q', func() {
			app.Stop()
		})

	content := tview.NewTextView().
		SetText("Welcome to GoalPost!").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	flex := tview.NewFlex().
		AddItem(menu, 0, 1, true).
		AddItem(content, 0, 4, false)

	menu.SetSelectedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		switch mainText {
		case "Fixtures":
			content.SetText("Fixtures")
		case "Results":
			content.SetText("Results")
		case "Tables":
			content.SetText("Tables")
		case "News":
			content.SetText("News")
		}
	})

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
