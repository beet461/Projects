package main

import (
	"log"
	"os"

	"github.com/asticode/go-astilectron"
)

func main() {
	// Initialize astilectron
	var a, _ = astilectron.New(log.New(os.Stderr, "", 0), astilectron.Options{
		AppName:            "Hi",
		VersionAstilectron: "0.33.0",
		VersionElectron:    "`6.1.2`",
	})
	defer a.Close()

	// Start astilectron
	a.Start()

	// Blocking
}
