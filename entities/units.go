package entities

import "time"

type Unit struct {
	Id        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
