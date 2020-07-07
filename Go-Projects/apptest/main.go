package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	c
)

func main() {
	a := app.New()
	w := a.NewWindow("Planet")

	w.SetContent(widget.NewLabel("Hello, Humans"))

	w.ShowAndRun()
}
