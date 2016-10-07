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

		if friends_user_1, ok := friends[user_id_1]; ok {
			var already_exist bool
			for _, f_id := range friends_user_1 {
				if f_id == user_id_2 {
					already_exist = true
					break
				}
			}
			if !already_exist {
				friends_user_1 = append(friends_user_1, user_id_2)
				friends[user_id_1] = friends_user_1
			}
		} else {
			friends[user_id_1] = []int64{user_id_2}
		}

		if friends_user_2, ok := friends[user_id_2]; ok {
			var already_exist bool
			for _, f_id := range friends_user_2 {
				if f_id == user_id_1 {
					already_exist = true
					break
				}
			}
			if !already_exist {
				friends_user_2 = append(friends_user_2, user_id_1)
				friends[user_id_2] = friends_user_2
			}
		} else {
			friends[user_id_2] = []int64{user_id_1}
		}
	}
	return nil
}

func add_friend() error {

	return nil
}
