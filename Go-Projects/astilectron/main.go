package main

import (
	"log"
	"os"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
)

func main() {
	// Initialize astilectron
	var a, _ = astilectron.New(log.New(os.Stderr, "", 0), astilectron.Options{
		AppName:            "uh",
		AppIconDefaultPath: "favicon.png", // If path is relative, it must be relative to the data director
		VersionAstilectron: "0.33.0",
		VersionElectron:    "6.1.2",
	})
	defer a.Close()

	// Start astilectron
	a.Start()

	var w, _ = a.NewWindow("index.html", &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(600),
		Width:  astikit.IntPtr(600),
	})

	w.Create()

	w.OpenDevTools()

	defer w.CloseDevTools()

	w.OnMessage(func(m *astilectron.EventMessage) interface{} {
		// Unmarshal
		var jsmsg string
		m.Unmarshal(&jsmsg)

		if jsmsg == "hello" {
			return "world"
		}
		return nil
	})

	// Blocking pattern
	a.Wait()
}
