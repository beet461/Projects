package main

import (
	"fmt"
	"os/exec"
	"log"
)

func main() {
	out, err := exec.Command("ls").Output()
	if err != nil {
		fmt.Println(err)
	}

	
}