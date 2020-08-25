package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("npm", "test").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out, err)
}
