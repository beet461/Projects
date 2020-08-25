package main

import "os/exec"

func main() {
	out, err := exec.Command("ls").Output()
}
