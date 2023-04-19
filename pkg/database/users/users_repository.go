package users

import (
	"database/sql"
	"errors"
	"fmt"
	"kaijuVpn/pkg/database"
	"log"
)

type User struct {
	ID          int
	Telegram_id string
	Is_active   bool
	Phone       sql.NullString
}

func getAllUsers(db database.Database) {
	rows, err := db.Connection.Query("select * from go_kaiju.products")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
}

func InsertUser(user User) (int64, error) {
	result, err := database.DB.Connection.Exec("INSERT INTO users (telegram_id) VALUES (?)", user.Telegram_id)
	if err != nil {
		log.Fatalln(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("InsertUser: %v", err)
	}
	return id, nil
}

func CreateIfNotExist(user User) (int64, error) {
	var db_user User

	row := database.DB.Connection.QueryRow("SELECT * FROM users WHERE telegram_id = ?", user.Telegram_id)
	err := row.Scan(&db_user.ID, &db_user.Telegram_id, &db_user.Phone, &db_user.Is_active)
	if err != nil && err != sql.ErrNoRows {
		log.Fatalln(err)
	}
	if err != nil && err == sql.ErrNoRows {
		return InsertUser(user)
	}
	return 0, errors.New("user already exist!")
}
