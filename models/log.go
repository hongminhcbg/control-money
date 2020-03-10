package models

import "time"

type Log struct {
	ID       int64      `gorm:"column:id;PRIMARY_KEY"`
	Detail   string     `gorm:"column:detail"`
	Money    int64      `gorm:"column:money"`
	Tag      string     `gorm:"column:tag"`
	UserID   int64      `gorm:"column:user_id"`
	CreateAt *time.Time `gorm:"column:created_at"`
	UpdateAt *time.Time `gorm:"column:updated_at"`
}
