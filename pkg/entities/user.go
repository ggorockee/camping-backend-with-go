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

type UserSerializer interface {
	TinyUserSerialize() TinyUserOutputSchema
	UserDetailSerialize() UserDetailOutputSchema
}

type userSerializer struct {
	User *User
}

func (u *userSerializer) TinyUserSerialize() TinyUserOutputSchema {
	return TinyUserOutputSchema{
		Id:       int(u.User.Id),
		Email:    u.User.Email,
		Username: u.User.Username,
		Role:     u.User.Role,
	}
}

func (u *userSerializer) UserDetailSerialize() UserDetailOutputSchema {
	return UserDetailOutputSchema{
		Id:       int(u.User.Id),
		Email:    u.User.Email,
		Username: u.User.Username,
		Role:     u.User.Role,
	}
}

func NewUserSerializer(u *User) UserSerializer {
	return &userSerializer{User: u}
}

// ============= input schema =============

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

// ============= output schema =============

type TinyUserOutputSchema struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type UserDetailOutputSchema struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
