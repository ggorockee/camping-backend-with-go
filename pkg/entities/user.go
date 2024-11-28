package entities

// User struct
type User struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Email    string `gorm:"uniqueIndex;not null;type:varchar(255)" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Username string `json:"username"`
}

type LoginInputSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpInputSchema struct {
	Email           string  `json:"email"`
	Password        string  `json:"password"`
	PasswordConfirm string  `json:"password_confirm"`
	Username        *string `json:"username"`
}

type ChangePasswordInputSchema struct {
	OldPassword        string `json:"old_password"`
	NewPassword        string `json:"new_password"`
	NewPasswordConfirm string `json:"new_password_confirm"`
}
