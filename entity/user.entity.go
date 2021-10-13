package entity

import "time"

type User struct {
	ID        int64  `gorm:"primary_key:auto_increment" json:"-"`
	Name      string `gorm:"type:varchar(100)" json:"username" validate:"required"`
	Email     string `gorm:"type:varchar(100);unique;" json:"email" validate:"required,email"`
	Password  string `gorm:"type:varchar(100)" json:"password" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
