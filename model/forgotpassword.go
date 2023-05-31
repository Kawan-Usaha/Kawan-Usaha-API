package Model

type ForgotPassword struct {
	//Email: The user's email that needs password reset.
	Email string `json:"email"`
}
