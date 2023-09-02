package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	gorm.Model

	Name string
}

type RoleResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Name      string    `json:"name"`
}
