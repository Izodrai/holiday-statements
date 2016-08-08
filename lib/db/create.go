package db

import (
	"io/ioutil"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
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