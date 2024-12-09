package entities

import (
	"time"
)

type Spot struct {
	Id         uint `json:"id" gorm:"primaryKey"`
	UserId     uint `json:"user_id"`
	User       User `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
	CategoryId *int `gorm:"default:null" json:"category_id"` // CategoryId가 null일 수가 있음
	// sqlite에서 SET NULL, mysql, postgresql에서는 SetNull
	// 배포시 아래 주석
	// sqlite 설정
	Category Category `gorm:"foreignKey:CategoryId;constraint:OnDelete:SET NULL;"`

	// 배포시 아래 주석해제
	// rds 설정
	//Category  Category  `gorm:"foreignKey:CategoryId;constraint:OnDelete:SetNull;"`
	Title     string    `json:"title"`
	Location  string    `json:"location"`
	Author    string    `json:"author"`
	Review    string    `json:"review"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	CoverImg string `json:"cover_img"`
}
