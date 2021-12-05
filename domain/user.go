package domain

type User struct {
	IIN       string `json:"iin"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}
