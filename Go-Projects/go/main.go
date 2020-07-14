package main

import "fmt"

type nums struct {
	on int
}

func main() {
	fmt.Println("start")
	var h int
	var hi nums
	fmt.Scan(&h)
	hi.on = h
	fmt.Println(hi.on)
}
