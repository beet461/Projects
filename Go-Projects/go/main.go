package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command(cmd/c, "ls").Output()
	if err != nil {
		fmt.Println(err)
	}
	output := string(out[:])
	fmt.Println(output)
	fmt.Println(out)
	fmt.Println(err)
}
