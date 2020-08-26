package main

import (
	"fmt"
	"os/exec"
	"os/user"
)

func main() {
	out, err := exec.Command("cmd", "/c", "dir").Output()
	me := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	output := string(out[:])
	fmt.Println(output)
	fmt.Println(me)
	//fmt.Println(out)
	//	fmt.Println(err)
}
