package entities

// User struct
type User struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Email    string `gorm:"uniqueIndex;not null;type:varchar(255)" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Username string `json:"username"`
}

type CreateUserSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}
