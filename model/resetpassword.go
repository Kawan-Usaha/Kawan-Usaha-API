package Model

type ResetPassword struct {
	VerificationCode string `json:"verification_code"`
	Password         string `json:"password"`
	PasswordConfirm  string `json:"password_confirm"`
}
