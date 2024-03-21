package dbase

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "dbname"
)

func ConnectToDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Открываем соединение с базой данных
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	// Проверяем, что соединение установлено
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected!")
	return db, nil
}
