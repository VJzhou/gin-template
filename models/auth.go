package models

type Auth struct {
	ID int `gorm:"primary_key json:"id""`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool{
	var auth Auth
	auth.Username = username
	auth.Password = password
	db.First(&auth)

	if auth.ID < 0 {
		return false
	}
	return true
}
