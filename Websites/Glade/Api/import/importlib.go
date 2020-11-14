package main

import (
	"fmt"
	"os/exec"
)

func importAll(rpstry string) {
	out, err := exec.Command("go", "get", "-u", rpstry).Output()
	if err != nil {
		fmt.Printf("Error importing : '%v'", rpstry)
	}
	output := string(out[:])
	fmt.Printf("Success importing : '%v' %v \n", rpstry, output)
}

func main() {
	rp1 := "github.com/gorilla/mux"
	importAll(rp1)
	rp2 := "github.com/lib/pq"
	importAll(rp2)
	rp3 := "github.com/rs/cors"
	importAll(rp3)
}
