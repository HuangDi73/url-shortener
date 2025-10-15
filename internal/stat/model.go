package stat

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Stat struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"create_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	LinkId    uint           `json:"link_id"`
	Clicks    int            `json:"clicks"`
	Date      datatypes.Date `json:"date"`
}
