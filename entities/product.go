package entities

import "time"

type Product struct {
	Id          uint
	Name        string
	Category    Category
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
