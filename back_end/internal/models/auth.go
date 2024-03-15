package models

type UserData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type User struct {
	Login    string `json:"auth_id"`
	Password string `json:"password_web"`
	Name     string `json:"name"`
	License  int    `json:"license"`
}

type HTTPError struct {
	ErrorCode    int
	ErrorMessage string
}
