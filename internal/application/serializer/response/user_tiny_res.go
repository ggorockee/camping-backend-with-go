package response

type UserTinyRes struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Email    string `gorm:"uniqueIndex;not null;type:varchar(255)" json:"email"`
	Username string `json:"username"`
	Role     string `json:"role" gorm:"default:'client'"`
}
