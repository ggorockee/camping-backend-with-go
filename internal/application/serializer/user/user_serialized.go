package userserializer

type TinyUserRes struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type DetailUserRes struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
