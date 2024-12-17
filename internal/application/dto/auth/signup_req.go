package authdto

type SignUpReq struct {
	Email           string  `json:"email"`
	Password        string  `json:"password"`
	PasswordConfirm string  `json:"password_confirm"`
	Username        *string `json:"username"`
}
