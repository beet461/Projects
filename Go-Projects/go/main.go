package main

import (
	"database/sql"
	"fmt"
)

func main() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "shashi3969"
		dbname   = "FileMaker"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

}
