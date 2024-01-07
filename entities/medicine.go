package entities

import (
	"time"
)

type Medicine struct {
	Id uint
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
}