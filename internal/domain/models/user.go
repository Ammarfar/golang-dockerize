package models

import (
	"time"
)

func (User) TableName() string {
	return "records"
}

func (UserWithTotalMarks) TableName() string {
	return "records"
}

type User struct {
	Id        *int64     `json:"id" gorm:"primaryKey"`
	Name      *string    `json:"name"`
	Marks     []Mark     `json:"marks"`
	CreatedAt *time.Time `json:"created_at"`
}

type UserWithTotalMarks struct {
	Id         *int64     `json:"id" gorm:"primaryKey"`
	CreatedAt  *time.Time `json:"createdAt"`
	TotalMarks *int       `json:"totalMarks"`
}

type UserFilter struct {
	StartDate *string
	EndDate   *string
	MinCount  *int
	MaxCount  *int
}
