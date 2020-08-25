package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("bash", "-c", "cd", "..").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Sprint(out, err)
}
