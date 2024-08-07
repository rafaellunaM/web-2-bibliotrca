package main

import (
	"database/sql"
	"fmt"
	"library/dbconfig"
)

var db *sql.DB
var err error

func checErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {

	fmt.Printf("Acessing %s...", dbconfig.DbName)

	db, err = sql.Open(dbconfig.PostgresDriver, dbconfig.DataSourceName)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("\nConnected")
	}

}
