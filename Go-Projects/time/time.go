package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	z, _ := t.Zone()
	fmt.Printf("\rThe time in your zone(%v) is %v", z, t)
	main()
}
