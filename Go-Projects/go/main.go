package main

import (
	"fmt"
	"os/exec"
	"os/user"
)

func main() {
	path := "~/Desktop/"
	cd, err := exec.Command("cmd", "/c", "cd", path).Output()
	out, err := exec.Command("cmd", "/c", "fsutil", "file", "createnew", "test", "0").Output()
	me, _ := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	output := string(out[:])
	cdout := string(cd[:])
	fmt.Println(cdout)
	fmt.Println(output)
	fmt.Println(me)
}
