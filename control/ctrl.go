package control

import (
	"database/sql"
)

func GetItems(db *sql.DB) ([]Item, error) {
	items := make([]Item, 0)
	rows, err := db.Query(`
		SELECT * FROM items
		ORDER BY id
	`)
	if err != nil {
		return []Item{}, err
	}
	for rows.Next() {
		var id int
		var name string
		var typee string
		var genre string
		err = rows.Scan(&id, &name, &typee, &genre)
		item := Item{Id: id, Name: name, Type: typee, Genre: genre}
		items = append(items, item)
	}
	return items, nil
}

func AddItem(item Item, db *sql.DB) error {
	_, err := db.Exec(`
		INSERT INTO items (name,type, genre)
		VALUES ( $1, $2, $3);
	`, item.Name, item.Type, item.Genre)
	if err != nil {
		return err
	}
	return nil
}

func UpdateItem(setArea string, whereArea string, value interface{}, neww string, db *sql.DB) error {
	_, err := db.Exec(`
		UPDATE items
		SET `+setArea+` = $1
		WHERE `+whereArea+` = $2;
	`, neww, value)
	if err != nil {
		return err
	}
	return nil
}

func DeleteItem(id int, db *sql.DB) error {
	_, err := db.Exec(`
		DELETE FROM items
		WHERE id = $1;
	`, id)
	if err != nil {
		return err
	}
	return nil
}
