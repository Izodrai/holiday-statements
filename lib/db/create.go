package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
)

func CreateDatabase(db *sql.DB) error {
	sqlb, err := ioutil.ReadFile("./db/create/create.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlb))
	if err != nil {
		return err
	}

	return nil
}
