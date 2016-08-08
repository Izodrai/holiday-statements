package db

import (
	"os"
	"errors"
	"strings"
	"../tools"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func Init() error {
	
	os.Remove("./db/save/spending.db") // to remove
	
	var db *sql.DB
	var err error
	var rows *sql.Rows
	
	db, err = sql.Open("sqlite3", "./db/save/spending.db")
	if err != nil {
		return err
	}
	defer db.Close()

	tools.GreenInfo("sqlite connected")
	
architecture_test:

	rows, err = db.Query("select id, login from users limit 1")
	if err != nil {
		if strings.Contains(err.Error(), "no such table: users") {
			tools.Info("sqlite empty, we need to feed it!")
			if err = CreateDatabase(db); err != nil {
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
		rows.Scan(&id,&login)
		
// 		tools.WhiteInfo("Login : ", login, " with id : ", id)
		
		if login != "vp" || id != "1" {
			return errors.New("bad values")
		}
	}
	
	tools.GreenInfo("sqlite ready")
	
	return nil
}