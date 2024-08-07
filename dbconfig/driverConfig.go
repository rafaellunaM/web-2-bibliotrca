package dbconfig

import "fmt"

type Books struct {
	ID    string
	Title string
	Body  []byte
}

const PostgresDriver = "postgres"

const User = "library"

const Password = "library"

const Host = "localhost"

const Port = "5432"

const DbName = "library"

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
