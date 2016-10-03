package db

import (
	"io/ioutil"
)

func CreateDatabase() error {
	sqlb, err := ioutil.ReadFile("./db/create/create.sql")
	if err != nil {
		return err
	}

	_, err = Db_connect.Exec(string(sqlb))
	if err != nil {
		return err
	}

	return nil
}
