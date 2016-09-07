package tools

func GenerateNav(login string) []string {
	nav := []string{
		"events",
	}
	
	if _, ok := Admins[login]; ok {
		nav = append(nav, "users")
	}
	
	return nav
}

func Unshift(arr []Spending, x Spending) []Spending {
	return append([]Spending{x}, arr...)
}