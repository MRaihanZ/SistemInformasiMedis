package entities

import "time"

type Pasien struct {
	Id        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}