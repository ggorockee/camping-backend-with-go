package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User struct
type User struct {
	Common
	Email    string `gorm:"uniqueIndex;not null;type:varchar(255)" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Username string `json:"username"`
	Role     string `json:"role" gorm:"default:'client'"`
}

func (s *User) GetId() string {
	return s.Id
}

func (s *User) BeforeCreate(tx *gorm.DB) (err error) {
	if s.GetId() == "" {
		id := uuid.New()
		s.Id = id.String()
	}
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
	return
}

func (s *User) BeforeSave(tx *gorm.DB) (err error) {
	s.UpdatedAt = time.Now()
	return
}
