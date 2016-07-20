package config

import (
	"encoding/json"
	"io/ioutil"
)

type User struct {
	User   string `json:"user"`
	Pwd    string `json:"pwd"`
	Rights int    `json:"rights"`
}

func LoadUsers(users map[string]User) error {
	
	usersConfigFile, err := ioutil.ReadFile("config/users.json")
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(usersConfigFile), &users); err != nil {
		return err
	}
	
	return nil
}