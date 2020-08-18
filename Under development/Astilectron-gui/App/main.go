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
		AppName:            "Hi",
		VersionAstilectron: "0.33.0",
		VersionElectron:    "`6.1.2`",
	})
	// Create a new window
	var w, _ = a.NewWindow("http://127.0.0.1:4000", &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(600),
		Width:  astikit.IntPtr(600),
	})

	defer a.Close()

	// Start astilectron
	a.Start()

	// Blocking
	a.Wait()
}
