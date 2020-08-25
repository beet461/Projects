package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("bash", "-c", "echo", "hello").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Sprint(out, err)
}
