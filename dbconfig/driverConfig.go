package dbconfig

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Book struct {
	ID           string  `json:"bookId"`
	Title        string  `json:"title"`
	Author       string  `json:"author"`
	Quantity     int     `json:"quantity"`
	Price        float64 `json:"price"`
	Availability bool    `json:"availability"`
}

var (
	PostgresDriver = "postgres"

	User     = getEnv("DB_USER", "library")
	Password = getEnv("DB_PASSWORD", "library")
	DbName   = getEnv("DB_NAME", "library")
	Host     = getEnv("DB_HOST", "localhost")
	Port     = getEnv("DB_PORT", "5432")

	TableName = "Book"
)

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

var DataSourceName = fmt.Sprintf(
	"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	Host, Port, User, Password, DbName,
)
