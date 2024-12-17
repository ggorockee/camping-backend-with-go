// migration 완료
package dto

type UserRole string

const (
	Client UserRole = "client"
	Admin  UserRole = "admin"
	Staff  UserRole = "staff"
	Owner  UserRole = "owner"
)

// ============= input schema =============

type LoginIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpIn struct {
	Email           string  `json:"email"`
	Password        string  `json:"password"`
	PasswordConfirm string  `json:"password_confirm"`
	Username        *string `json:"username"`
}

type ChangePasswordIn struct {
	OldPassword        string `json:"old_password"`
	NewPassword        string `json:"new_password"`
	NewPasswordConfirm string `json:"new_password_confirm"`
}

// ============= output schema =============

type TinyUserOut struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type UserDetailOut struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
