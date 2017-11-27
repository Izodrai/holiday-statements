package db

import (
	"../tools"
	"database/sql"
	"errors"
	"github.com/izodrai/utils/logs"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"strings"
)

var Db_connect *sql.DB

func Init_db_connect() error {

	//////////////////////////////////
	// to remove
	os.Remove("./db/save/spending.db")
	//////////////////////////////////

	var err error
	var rows *sql.Rows

	Db_connect, err = sql.Open("sqlite3", "./db/save/spending.db")
	if err != nil {
		return err
	}

	logs.Green_info("sqlite connected")

architecture_test:

	rows, err = Db_connect.Query("select id, login from users limit 1")
	if err != nil {
		if strings.Contains(err.Error(), "no such table: users") {
			logs.Info("sqlite empty, we need to feed it!")
			if err = CreateDatabase(); err != nil {
				return err
			}
			logs.Info("sqlite generated")
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

	logs.Green_info("sqlite ready")

	return nil
}

func Init_system() error {

	tools.Users = make(map[string]tools.User)
	tools.Users_id = make(map[int64]tools.User)
	tools.Connected_users = make(map[int64]tools.User)
	tools.Admins = make(map[string]tools.User)
	tools.Admins_id = make(map[int64]tools.User)
	tools.Friends = make(map[int64][]int64)

	if err := load_users(tools.Users); err != nil {
		return err
	}

	for login, user := range tools.Users {
		tools.Users_id[user.Id] = user
		if user.Admin {
			tools.Admins[login] = user
			tools.Admins_id[user.Id] = user
		}
	}

	if err := load_friends(tools.Friends); err != nil {
		return err
	}
	return nil
}
