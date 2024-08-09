package main

import (
	"fmt"
	"library/tables"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func checErr(err error) {
	if err != nil {
		log.Fatalf("Erro: %v", err)
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	host := os.Getenv("db_host")
	user := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	db_name := os.Getenv("db_name")
	db_port := os.Getenv("db_port")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Recife",
		host, user, password, db_name, db_port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	checErr(err)

	db.AutoMigrate(&tables.Librarian{}, &tables.Admin{}, &tables.Client{}, &tables.Book{}, &tables.Loan{}, &tables.Fine{})

	var loans []tables.Loan

	err = db.Preload("Client").Preload("Book").Find(&loans).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, loan := range loans {
		fmt.Printf("LoanID: %s, CPF: %s, BookID: %s, DueDate: %s, ReturnDate: %v, ClientName: %s, BookTitle: %s\n",
			loan.LoanID, loan.CPF, loan.BookID, loan.DueDate.Format("2006-01-02"), loan.ReturnDate, loan.Client.Name, loan.Book.Title)
	}
}
