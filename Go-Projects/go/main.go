package main

import (
	"fmt"
	"os/exec"
)

func useless(out, err) {
	return out, err
}

func main() {
	out, err := exec.Command("cmd", "ls").Output()
	if err != nil {
		fmt.Println(err)
	}
	useless(out, err)
}
