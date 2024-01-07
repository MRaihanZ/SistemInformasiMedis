package entities

import "time"

type Doctor struct {
	Id         uint
	Name       string
	CreatedAt time.Time
	UpdatedAt time.Time
}