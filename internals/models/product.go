package models

import (
	"time"

	"gorm.io/gorm"
)

type (
	Product struct {
		ID        string         `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
		CreatedAt time.Time      `json:"created_at"`
		UpdatedAt time.Time      `json:"updated_at"`
		DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	}
)
