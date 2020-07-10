package main

import (
	"fmt"
	"time"
)

func main() {
	var c float64
	var entry string
	c = 0
	h := false
	fmt.Println("Would you like a one, microsecond, millisecond, second, minute, hour, day or year counter?")
	fmt.Scan(&entry)
	switch entry {
	case "microsecond":
		for !h {
			c++
			time.Sleep(1 * time.Microsecond)
			fmt.Println(c)
		}
	case "millisecond":
		for !h {
			c++
			time.Sleep(1 * time.Millisecond)
			fmt.Println(c)
		}
	case "second":
		for !h {
			c++
			time.Sleep(1 * time.Second)
			fmt.Println(c)
		}
	case "minute":
		for !h {
			c++
			time.Sleep(1 * time.Minute)
			fmt.Println(c)
		}
	case "hour":
		for !h {
			c++
			time.Sleep(1 * time.Hour)
			fmt.Println(c)
		}
	case "day":
		for !h {
			c++
			time.Sleep(24 * time.Hour)
			fmt.Println(c)
		}
	case "year":
		for !h {
			c++
			time.Sleep(8760 * time.Hour)
			fmt.Println(c)
		}
	}

}
