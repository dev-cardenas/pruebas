package domain_user

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}
