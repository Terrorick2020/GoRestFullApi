package internal

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name" binding:"*required"`
	Username string `json:"username" binding:"*required"`
	Password string `json:"password" binding:"required"`
}
