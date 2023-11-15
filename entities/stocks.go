package entities

import "time"

type Stock struct {
	Id        uint
	Name      string
	UpdatedAt time.Time
	CreatedAt time.Time
}
