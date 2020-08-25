package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("go", "run", "index.js").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("hello")
	fmt.Println(out, err)
}
