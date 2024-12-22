package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Amenity struct {
	Common
	Name        string  `json:"name" gorm:"type:varchar(20);"`
	Description *string `json:"description"`
}

func (s *Amenity) GetId() string {
	return s.Id
}

func (s *Amenity) IsExist() bool {
	return s.Id != ""
}

func (s *Amenity) BeforeCreate(tx *gorm.DB) (err error) {
	if s.GetId() == "" {
		id := uuid.New()
		s.Id = id.String()
	}
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
	return
}

func (s *Amenity) BeforeSave(tx *gorm.DB) (err error) {
	s.UpdatedAt = time.Now()
	return
}
