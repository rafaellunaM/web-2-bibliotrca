package tables

import (
	"time"
)

type Librarian struct {
	CPF         string  `gorm:"primaryKey;column:cpf"`
	Name        string  `gorm:"column:name"`
	Email       *string `gorm:"column:email"`
	Password    string  `gorm:"column:password"`
	PhoneNumber *string `gorm:"column:phone_number"`
}

func (Librarian) TableName() string {
	return "librarian"
}

type Admin struct {
	CPF         string  `gorm:"primaryKey;column:cpf"`
	Name        string  `gorm:"column:name"`
	Email       *string `gorm:"column:email"`
	Password    string  `gorm:"column:password"`
	PhoneNumber *string `gorm:"column:phone_number"`
}

func (Admin) TableName() string {
	return "admin"
}

type Client struct {
	CPF         string  `gorm:"primaryKey;column:cpf"`
	Name        string  `gorm:"column:name"`
	Email       *string `gorm:"column:email"`
	Password    string  `gorm:"column:password"`
	PhoneNumber *string `gorm:"column:phone_number"`
}

func (Client) TableName() string {
	return "client"
}

type Book struct {
	BookID       string `gorm:"primaryKey;column:bookId"`
	Title        string `gorm:"column:title"`
	Author       string `gorm:"column:author"`
	Category     string `gorm:"column:category"`
	Availability bool   `gorm:"column:availability"`
}

func (Book) TableName() string {
	return "book"
}

type Loan struct {
	LoanID     string     `gorm:"column:loanId;primaryKey"`
	CPF        string     `gorm:"column:cpf"`
	BookID     string     `gorm:"column:bookId"`
	DueDate    time.Time  `gorm:"column:dueDate"`
	ReturnDate *time.Time `gorm:"column:returnDate"`
	Client     Client     `gorm:"foreignKey:CPF;references:CPF"`
	Book       Book       `gorm:"foreignKey:BookID;references:BookID"`
}

func (Loan) TableName() string {
	return "loan"
}

type LoanDetail struct {
	LoanID     string     `gorm:"column:loanId"`
	CPF        string     `gorm:"column:cpf"`
	BookID     string     `gorm:"column:bookId"`
	DueDate    *time.Time `gorm:"column:dueDate"`
	ReturnDate time.Time  `gorm:"column:returnDate"`
	ClientName string     `gorm:"column:clientName"`
	BookTitle  string     `gorm:"column:bookTitle"`
}

type Fine struct {
	FineID string  `gorm:"column:fineId;primaryKey"`
	CPF    string  `gorm:"column:cpf"`
	LoanID string  `gorm:"column:loanId"`
	Amount float64 `gorm:"column:amount"`
	Paid   bool    `gorm:"column:paid"`
	Client Client  `gorm:"foreignKey:CPF;references:CPF"`
	Loan   Loan    `gorm:"foreignKey:LoanID;references:LoanID"`
}

func (Fine) TableName() string {
	return "fine"
}
