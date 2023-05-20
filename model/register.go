package Model

type Register struct {
	//Name: The user's name.
	//Email: The user's email.
	//Password: The user's password.
	//PasswordConfirm: The user's password confirmation.
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}
