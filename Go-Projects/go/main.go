package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	defer db.Close()
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

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	sqlStatement := `
	INSERT INTO api_keys (user, api_key)
	VALUES ('TeSt', 'keytimeemeememem');`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}

}
