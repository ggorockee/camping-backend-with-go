package entities

// User struct
type User struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Email    string `gorm:"uniqueIndex;not null;type:varchar(255)" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Username string `json:"username"`
	Role     string `json:"role" gorm:"default:'client'"`
}

// like enum
type UserRole string

const (
	Client UserRole = "client"
	Owner  UserRole = "owner"
	Staff  UserRole = "staff"
	Admin  UserRole = "admin"
)

func (u *User) TinyUserSerialize() TinyUserOutputSchema {
	return TinyUserOutputSchema{
		Id:       u.Id,
		Email:    u.Email,
		Username: u.Username,
		Role:     u.Role,
	}
}

func (u *User) UserDetailSerialize() UserDetailOutputSchema {
	return UserDetailOutputSchema{
		Id:       u.Id,
		Email:    u.Email,
		Username: u.Username,
		Role:     u.Role,
	}
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

type TinyUserOutputSchema struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type UserDetailOutputSchema struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
