package main

import (
	"fmt"
	"os/exec"
	"os/user"
)

func main() {
	out, err := exec.Command("cmd", "/c", "dir").Output()
	me, m := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	output := string(out[:])
	fmt.Println(output)
	fmt.Println(me)
}
