package main

import (
	"fmt"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func basic() string {
	h := fmt.Sprintln("Hello ")
	return h
}

func main() {

	js := mewn.String("./frontend/build/main.js")
	css := mewn.String("./frontend/build/main.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "wailapp",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(basic)
	app.Run()
}
