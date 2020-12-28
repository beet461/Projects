package gladeapi

import (
	"database/sql"
	"fmt"
)

func connect() *sql.DB {
	//Defining the values for connecting to the database
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "shashi3969"
		dbname   = "Glade"
	)

	//Database values
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	//Connecting/opening to the database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	//Closing the database at the end
	defer db.Close()

	//Pinging the database for any errors
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
