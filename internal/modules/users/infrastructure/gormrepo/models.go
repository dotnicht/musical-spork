package gormrepo

import "time"

type userModel struct {
	ID        string    `gorm:"type:uuid;primaryKey"`
	Email     string    `gorm:"type:text;not null;uniqueIndex"`
	Name      string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (userModel) TableName() string { return "users" }
