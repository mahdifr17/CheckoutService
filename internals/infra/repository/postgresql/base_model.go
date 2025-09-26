package postgresql

import (
	"time"

	"gorm.io/gorm"
)

type (
	BaseModel struct {
		ID        string         `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
		CreatedAt time.Time      `gorm:"autoCreateTime"`
		UpdatedAt time.Time      `gorm:"autoUpdateTime"`
		DeletedAt gorm.DeletedAt `gorm:"index"`
	}
)
