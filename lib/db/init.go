package db

import (
	"../tools"
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"strings"
)

var Db_connect *sql.DB

func Init_connect_and_db() error {

	os.Remove("./db/save/spending.db") // to remove

	var err error
	var rows *sql.Rows

	Db_connect, err = sql.Open("sqlite3", "./db/save/spending.db")
	if err != nil {
		return err
	}

	tools.Green_info("sqlite connected")

architecture_test:

	rows, err = Db_connect.Query("select id, login from users limit 1")
	if err != nil {
		if strings.Contains(err.Error(), "no such table: users") {
			tools.Info("sqlite empty, we need to feed it!")
			if err = CreateDatabase(); err != nil {
				return err
			}
			tools.Info("sqlite generated")
			goto architecture_test
		} else {
			return err
		}
	}
	defer rows.Close()

	for rows.Next() {
		var id, login string
		rows.Scan(&id, &login)

		// 		tools.White_info("Login : ", login, " with id : ", id)

		if login != "admin" || id != "1" {
			return errors.New("bad values")
		}
	}

	tools.Green_info("sqlite ready")

	return nil
}