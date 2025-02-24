package internal

type User struct {
	ID       int    `json:"_" db:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
