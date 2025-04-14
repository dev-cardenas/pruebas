package application_user

type CreateUserDTO struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}
