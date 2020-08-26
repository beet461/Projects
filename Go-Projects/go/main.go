package main

import (
	"fmt"
	"os/exec"
	"os/user"
)

func main() {
	out, err := exec.Command("cmd", "/c", "fsutil", "file", "createnew", "test", "0").Output()
	me, _ := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	output := string(out[:])
	fmt.Sprint(output)
	fmt.Println(me)
}
