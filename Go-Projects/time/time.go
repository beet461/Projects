package main

import (
	"fmt"
	"time"
)

func main() {
	h := "h"
	for h == "h" {
		t := time.Now()
		z, _ := t.Zone()
		fmt.Printf("\rThe time in your zone(%v) is %v", z, t)
	}
}
