package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Planet-BH564VN89")

	w.SetContent(widget.NewLabel("Hello, Humans"))

	w.ShowAndRun()
}
