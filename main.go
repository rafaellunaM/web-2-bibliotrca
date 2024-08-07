package main

import (
	"fmt"
	"library/tables"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

func checErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	dsn := "host=localhost user=library password=library dbname=library port=5432 sslmode=disable TimeZone=America/Recife"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	checErr(err)

	err = db.AutoMigrate(&tables.Librarian{})
	checErr(err)

	if err != nil {
		panic("connection fail")
	}

	db.AutoMigrate(&tables.Librarian{})

	var librarian tables.Librarian

	result := db.First(&librarian)
	if result.Error != nil {
		fmt.Println("Error:", result.Error)
	} else {
		fmt.Printf("Librarian: %+v\n", librarian)
	}
}
