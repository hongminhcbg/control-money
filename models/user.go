package models

import "time"

type User struct {
	ID       int64      `gorm:"column:id;PRIMARY_KEY"`
	Name     string     `gorm:"column:name"`
	Username string     `gorm:"column:username"`
	Money    int64      `gorm:"column:money"`
	Password string     `gorm:"column:password"`
	CreateAt *time.Time `gorm:"column:created_at"`
	UpdateAt *time.Time `gorm:"column:updated_at"`
}
