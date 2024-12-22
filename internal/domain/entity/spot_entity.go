package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Spot struct {
	Common
	UserId string `json:"user_id" gorm:"type:varchar(255)"`
	User   User   `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`

	Name        string `json:"name"`
	Country     string `json:"country"`
	City        string `json:"city"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Address     string `json:"address"`
	PetFriendly bool   `json:"pet_friendly"`

	CategoryId *string  `gorm:"type:varchar(255);default:null" json:"category_id"` // CategoryId가 null일 수가 있음
	Category   Category `gorm:"foreignKey:CategoryId;constraint:OnDelete:SET NULL;"`

	Amenities []Amenity `gorm:"many2many:spot_amenities"`

	Reviews []Review `gorm:"foreignKey:SpotId" json:"reviews"`
}

func (s *Spot) GetId() string {
	return s.Id
}

func (s *Spot) BeforeCreate(tx *gorm.DB) (err error) {
	if s.GetId() == "" {
		id := uuid.New()
		s.Id = id.String()
	}

	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
	return
}

func (s *Spot) BeforeSave(tx *gorm.DB) (err error) {
	s.UpdatedAt = time.Now()
	return
}
