package main

import (
	"fmt"
	"log"
)

func queryDb(queryStatement string) (scanValues [3]string) {
	db := connect()

	queryResult, err := db.Query(queryStatement)
	if err != nil {
		fmt.Println("Query error : ", err)
	}

	var scanResult [3]string

	for queryResult.Next() {
		if err := queryResult.Scan(&scanResult[0], &scanResult[1], &scanResult[2]); err != nil {
			log.Fatal(err)
		}
	}

	return scanResult
}
