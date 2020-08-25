package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command(["cd", ".."]).Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out, err)
}
