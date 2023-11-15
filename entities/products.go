package entities

import "time"

type Product struct {
	Id           uint
	Name         string
	Category     Category
	Stock        int64
	Description  string
	Serialnumber string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
