package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Connection *sql.DB
}

var DB = Database{}

func(d *Database) Connect ()  {
	conn, err := sql.Open("mysql", os.Getenv("MYSQL_DSN"))
	if err != nil {
		log.Fatalln(err)
	}
	d.Connection = conn
}