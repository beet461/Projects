package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("cmd", "echo", "hello").Output()
	if err != nil {
		fmt.Println(err)
	}
	output := string(out[:])
	fmt.Println(output)
}
