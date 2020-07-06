package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Does this work?")

	w.SetContent(widget.NewLabel("Hello, Humans"))

	w.ShowAndRun()
}
