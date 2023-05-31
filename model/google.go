package Model

type GoogleLogin struct {
	//Sub: The user's unique Google ID.
	//Name: The user's name.
	//GivenName: The user's given name.
	//FamilyName: The user's family name.
	//Email: The user's email.
	//Verified: Whether the user's email is verified.
	//Picture: The user's profile picture.
	//Locale: The user's locale.
	Sub        string `json:"sub"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Email      string `json:"email"`
	Verified   bool   `json:"verified_email"`
	Picture    string `json:"picture"`
	Locale     string `json:"locale"`
}
