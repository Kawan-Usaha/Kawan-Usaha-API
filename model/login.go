package Model

type Login struct {
	//Email: The user's email.
	//Password: The user's password.
	Email    string `json:"email"`
	Password string `json:"password"`
}
