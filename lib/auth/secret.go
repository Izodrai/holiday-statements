package auth

import (
	"../db"
	"../tools"
)

func Init() error {

	tools.Users = make(map[string]tools.User)
	tools.UsersId = make(map[int64]tools.User)
	tools.Admins = make(map[string]tools.User)

	if err := db.LoadUsers(tools.Users); err != nil {
		return err
	}

	for login, user := range tools.Users {
		tools.UsersId[user.Id]= user
		if user.Admin {
			tools.Admins[login] = user
		}
	}

	return nil
}

func UserSecret(login, realm string) string {

	if validUser, ok := tools.Users[login]; ok {
		return validUser.Password
	}
	return ""
}

func AdminSecret(login, realm string) string {

	if validAdmin, ok := tools.Admins[login]; ok {
		return validAdmin.Password
	}
	return ""
}

func UpdateSecret(user tools.User) {
	tools.Users[user.Login] = user
	tools.UsersId[user.Id]= user
	for login, user := range tools.Users {
		if user.Admin {
			tools.Admins[login] = user
		}
	}
}