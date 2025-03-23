package model

import (
	"time"
)

// User は複数の Todo を持つ
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:100;uniqueIndex;not null"`
	Password  string `gorm:"size:255;not null"`
	Todos     []Todo `gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Todo は 1 User に属し、複数の Memo を持つ
type Todo struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"index;not null"`
	Title       string `gorm:"size:200;not null"`
	Description string `gorm:"type:text"`
	Completed   bool   `gorm:"default:false"`
	Deadline    *time.Time
	Memos       []Memo `gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	User User `gorm:"foreignKey:UserID"`
}

// Memo は 1 Todo に属する
type Memo struct {
	ID        uint   `gorm:"primaryKey"`
	TodoID    uint   `gorm:"index;not null"`
	Content   string `gorm:"type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Todo Todo `gorm:"foreignKey:TodoID"`
}
