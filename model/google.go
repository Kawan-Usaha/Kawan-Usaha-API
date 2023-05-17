package Model

type GoogleLogin struct {
	Sub        string `json:"sub"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Email      string `json:"email"`
	Verified   bool   `json:"verified_email"`
	Picture    string `json:"picture"`
	Locale     string `json:"locale"`
}
