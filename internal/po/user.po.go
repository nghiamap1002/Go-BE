package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"column:uuid; type:varchar(255); not null; index:idx_uuid: unique;"`
	Username string    `gorm:"column:username"`
	IsActive bool      `gorm:"column:is_active; type:boolean;"`
	Roles    []Role    `gorm:"many2many:user_roles"`
}

// func (User)
func (u *User) TableName() string {
	return "user"
}
