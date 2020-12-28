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
	const (
		mux = "github.com/gorilla/mux"
		lib = "github.com/lib/pq"
		rscors = "github.com/rs/cors"
		respond = "gopkg.in/matryer/respond.v1"
	)

	importAll(mux)
	importAll(lib)
	importAll(rscors)
	importAll(respond)
}
