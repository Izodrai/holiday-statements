package authentification

import (
	"../db"
	"../tools"
)

func Init_authenfication() error {

	tools.Users = make(map[string]tools.User)
	tools.Users_id = make(map[int64]tools.User)
	tools.Connected_users = make(map[int64]tools.User)
	tools.Admins = make(map[string]tools.User)
	
	if err := db.Load_users(tools.Users); err != nil {
		return err
	}

	for login, user := range tools.Users {
		tools.Users_id[user.Id]= user
		if user.Admin {
			tools.Admins[login] = user
		}
	}

	return nil
}