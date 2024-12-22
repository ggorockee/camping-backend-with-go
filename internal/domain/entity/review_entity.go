package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	Common
	UserId string `json:"user_id" gorm:"type:varchar(255)"`
	User   User   `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`

	SpotId *string `json:"spot_id" gorm:"type:varchar(255)"`
	Spot   Spot    `gorm:"foreignKey:SpotId;constraint:OnDelete:SET NULL"`

	Payload string `json:"payload" gorm:"type:text"`
	Rating  int    `json:"rating"`
}

func (s *Review) GetId() string {
	return s.Id
}

func (s *Review) BeforeCreate(tx *gorm.DB) (err error) {
	if s.GetId() == "" {
		id := uuid.New()
		s.Id = id.String()
	}
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
	return
}

func (s *Review) BeforeSave(tx *gorm.DB) (err error) {
	s.UpdatedAt = time.Now()
	return
}
