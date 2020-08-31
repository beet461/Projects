package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	rand.Seed(time.Now().UnixNano())
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
	username := "yeEYeeyyEryr"
	key := rand.Intn(100000)
	fmt.Println(username, key)
	sqlStatement := fmt.Sprintf(`
	INSERT INTO api_keys (username, api_key)
	VALUES ('%v', '%v');`, username, key)
	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}

}
