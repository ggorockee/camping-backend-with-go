package amenityentity

import "time"

type Amenity struct {
	Id          uint      `json:"id" gorm:"primary_key"`
	Name        string    `json:"name" gorm:"type:varchar(20);"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
