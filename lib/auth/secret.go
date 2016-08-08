package auth

import (
	"../tools"
)

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