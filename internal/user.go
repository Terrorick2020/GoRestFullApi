package internal

// User represents a user in the system
// swagger:model
type User struct {
	ID       int    `json:"_" db:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
