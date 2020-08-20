package main

import (
	"log"
	"os"

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

	// Blocking pattern
	a.Wait()
}
