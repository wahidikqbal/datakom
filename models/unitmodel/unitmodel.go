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
	result, err := config.DB.Exec(`INSERT INTO units (name, created_at, updated_at) VALUE (?, ?, ?)`, unit.Name, unit.CreatedAt, unit.UpdatedAt)
	if err != nil {
		panic(err)
	}

	lastinsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastinsertId > 0

}

func Detail(id int) entities.Unit {
	row := config.DB.QueryRow(`SELECT id, name FROM units WHERE id = ?`, id)

	var unit entities.Unit
	if err := row.Scan(&unit.Id, &unit.Name); err != nil {
		panic(err)
	}

	return unit

}

func Update(id int, unit entities.Unit) bool {
	query, err := config.DB.Exec(`UPDATE units SET name = ?, updated_at = ? WHERE id = ?`, unit.Name, unit.UpdatedAt, id)
	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec(`DELETE FROM units WHERE id = ?`, id)
	return err
}
