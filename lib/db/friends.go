package db

import (
	"database/sql"
)

func load_friends(friends map[int64][]int64) error {
	var err error
	var rows *sql.Rows

	rows, err = Db_connect.Query(`select user_id_1, user_id_2 from friends`)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {

		var user_id_1, user_id_2 int64

		err = rows.Scan(&user_id_1, &user_id_2)
		if err != nil {
			return err
		}
		
		if val, ok := friends[user_id_1]; ok {
			val = append(val, user_id_2)
			friends[user_id_1] = val
		} else {
			friends[user_id_1] = []int64{user_id_2}
		}
	}
	return nil
}
