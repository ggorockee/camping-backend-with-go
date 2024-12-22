package entity

import "time"

type Common struct {
	Id        string    `json:"id" gorm:"primaryKey;type:varchar(50)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
