package models

import "time"

type Log struct {
	ID       int64 `gorm:"column:id;PRIMARY_KEY"`
	Detail   int64 `gorm:"column:detail"`
	Money    int64 `gorm:"column:money"`
	Tag      int64 `gorm:"column:tag"`
	UserID   int64 `gorm:"column:user_id"`
	CreateAt *time.Time
	UpdateAt *time.Time
}
