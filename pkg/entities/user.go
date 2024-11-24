package entities

// User struct
type User struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Username string `json:"username"`
}
