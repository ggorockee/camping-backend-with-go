package entity

// User struct
type User struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Email    string `gorm:"uniqueIndex;not null;type:varchar(255)" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Username string `json:"username"`
	Role     string `json:"role" gorm:"default:'client'"`
}
