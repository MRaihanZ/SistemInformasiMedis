package entities

import "time"

type Record struct {
	Id          uint
	Pasien Pasien
	Category    Category
	Doctor Doctor
	Diagnose string
	Description string
	Medicine Medicine
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
