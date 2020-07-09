package main

import (
	"log"
	"os"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
)

func main() {
	var a, _ = astilectron.New(log.New(os.Stderr, "", 0), astilectron.Options{
		AppName:            "Hello",
		AppIconDefaultPath: "icon.png",
		VersionAstilectron: "0.33.0",
		VersionElectron:    "6.1.2",
	})
	defer a.Close()

	a.Start()

	var w, _ = a.NewWindow("index.html", &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(600),
		Width:  astikit.IntPtr(600),
	})

	w.Create()

	a.Wait()
}
