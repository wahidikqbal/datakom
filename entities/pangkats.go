package entities

import "time"

type Pangkat struct {
	Id        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
