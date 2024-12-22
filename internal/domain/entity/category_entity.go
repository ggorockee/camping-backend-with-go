package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	Common
	Name  string `json:"name" gorm:"uniqueIndex;type:varchar(20)"`
	Spots []Spot `json:"spot"`
}

func (s *Category) GetId() string {
	return s.Id
}

func (s *Category) BeforeCreate(tx *gorm.DB) (err error) {
	if s.GetId() == "" {
		id := uuid.New()
		s.Id = id.String()
	}
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
	return
}

func (s *Category) BeforeSave(tx *gorm.DB) (err error) {
	s.UpdatedAt = time.Now()
	return
}

func (s *Category) IsExist() bool {
	return s.Id != ""
}

func (c *Category) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}{
		Id:   c.Id,
		Name: c.Name,
	})
}
