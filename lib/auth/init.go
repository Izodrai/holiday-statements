package auth

import (
	"../tools"
	"io/ioutil"
	"encoding/json"
)

func Init() error {
	
	tools.Users = make(map[string]tools.User)
	tools.Admins = make(map[string]tools.User)
	
	if err := loadUsers(tools.Users); err != nil {
		return err
	}
	
	for login, user := range tools.Users {
		if user.Admin {
			tools.Admins[login]=user
		}
	}
	
	return nil
}

func loadUsers(users map[string]tools.User) error {
	
	file, err := ioutil.ReadFile("config/users.json")
	if err != nil {
		return err
	}
	
	if err := json.Unmarshal([]byte(file), &users); err != nil {
		return err
	}
	
	return nil
}