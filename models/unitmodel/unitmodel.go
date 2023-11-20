package unitmodel

import (
	"go-web/config"
	"go-web/entities"
)

func GetAll() []entities.Unit {
	rows, err := config.DB.Query(`SELECT * FROM units`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var units []entities.Unit

	for rows.Next() {
		var unit entities.Unit

		if err := rows.Scan(&unit.Id, &unit.Name, &unit.CreatedAt, &unit.UpdatedAt); err != nil {
			panic(err)
		}

		units = append(units, unit)
	}
	return units

}

func Create(unit entities.Unit) bool {
	result, err := config.DB.Exec(`INSERT INTO table units (name, created_at, updated_at), VALUE (?, ?, ?)`, unit.Name, unit.CreatedAt, unit.UpdatedAt)
	if err != nil {
		panic(err)
	}

	lastinsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastinsertId > 0
}
