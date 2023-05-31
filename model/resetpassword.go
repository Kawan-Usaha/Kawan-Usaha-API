package Model

type ResetPassword struct {
	//VerificationCode: The verification code sent to the user's email.
	//PasswordConfirm: The new password confirmation.
	//Password: The new password.
	VerificationCode string `json:"verification_code"`
	Password         string `json:"password"`
	PasswordConfirm  string `json:"password_confirm"`
}
