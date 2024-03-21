package control

import (
	"database/sql"
	"fmt"
)

type Item struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"` // movie / series / cartoon
	Genre string `json:"genre"`
}

func CreateItemsTable(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS items(
	    id SERIAL PRIMARY KEY,
	    name VARCHAR(255) NOT NULL,
	    type VARCHAR(255),
	    genre VARCHAR(255)
	);
	`)
	if err != nil {
		return err
	}
	fmt.Println("Table successfuly created")
	return nil
}
