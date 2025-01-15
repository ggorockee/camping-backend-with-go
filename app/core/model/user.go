package model

type User struct {
	Id       string  `json:"id" gorm:"primaryKey"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	UserName *string `json:"user_name"`
}
