package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
)

func main() {
	var yesno string
	path := "~/Desktop/Projects/Go-Projects/go/"
	exec.Command("cmd", "/c", "cd", path).Output()
	exec.Command("cmd", "/c", "fsutil", "file", "createnew", "test", "0")
	me, _ := user.Current()
	fmt.Println(me)
	fmt.Println("Would you like to delete file? (yes or no)")
	fmt.Scan(&yesno)
	if yesno == "yes" {
		exec.Command("cmd", "/c", "del", "test")
	} else {
		os.Exit(200)
	}
}
