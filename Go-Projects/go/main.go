package main

import (
	"fmt"
	"os/exec"
	"os/user"
)

func main() {
	//path := "~/Desktop/Projects/Go-Projects/go/"
	//	exec.Command("cmd", "/c", "cd", path).Output()
	//	out, err := exec.Command("cmd", "/c", "fsutil", "file", "createnew", "test", "0").Output()
	out, err := exec.Command("git", "ls").Output()
	me, _ := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	output := string(out[:])
	//cdout := string(cd[:])
	//fmt.Println(cdout)
	fmt.Println(output)
	fmt.Println(me)
}
