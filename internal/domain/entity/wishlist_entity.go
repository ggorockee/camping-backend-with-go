package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WishList struct {
	Common
	Name string `json:"name" gorm:"type:varchar(150)"`

	//Spots     interfaces.SpotCollection `gorm:"-"`                       // GORM에서 무시
	Spots []Spot `gorm:"many2many:wishlist_spot"` // 실제 DB 관계

	UserId string `json:"user_id" gorm:"type:varchar(255)"`
	User   User   `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
}

func (s *WishList) GetId() string {
	return s.Id
}

func (s *WishList) BeforeCreate(tx *gorm.DB) (err error) {
	if s.GetId() == "" {
		id := uuid.New()
		s.Id = id.String()
	}
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
	return
}

func (s *WishList) BeforeSave(tx *gorm.DB) (err error) {
	s.UpdatedAt = time.Now()
	return
}
