package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("go", "run", "main2.go").Output()
	if err != nil {
		fmt.Println(err)
	} else {
		done()
	}
	fmt.Println(out, err)
}
