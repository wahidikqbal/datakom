package entities

import "time"

type Product struct {
	Id           uint
	Name         string
	Pangkat      Pangkat
	Nrp          string
	Kesatuan     Kesatuan
	Category     Category
	Serialnumber string
	Stock        Stock
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
