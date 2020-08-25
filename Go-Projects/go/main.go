package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("ls").Output()
	fmt.Println(out, err)
}
