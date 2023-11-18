package pangkatmodel

import (
	"go-web/config"
	"go-web/entities"
)

func GetAll() []entities.Pangkat {
	rows, err := config.DB.Query("SELECT * FROM pangkats")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var pangkats []entities.Pangkat

	for rows.Next() {
		var pangkat entities.Pangkat
		if err := rows.Scan(&pangkat.Id, &pangkat.Name, &pangkat.CreatedAt, &pangkat.UpdatedAt); err != nil {
			panic(err)
		}

		pangkats = append(pangkats, pangkat)
	}

	return pangkats
}

func Create(pangkat entities.Pangkat) bool {
	result, err := config.DB.Exec(`INSERT INTO pangkats (name, created_at, updated_at)
	VALUE (?, ?, ?)`, pangkat.Name, pangkat.CreatedAt, pangkat.UpdatedAt)
	if err != nil {
		panic(err)
	}

	lastinsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastinsertId > 0
}

func Detail(id int) entities.Pangkat {
	row := config.DB.QueryRow(`SELECT id, name FROM pangkats WHERE id = ?`, id)

	var pangkat entities.Pangkat
	if err := row.Scan(&pangkat.Id, &pangkat.Name); err != nil {
		panic(err)
	}

	return pangkat
}

func Update(id int, pangkat entities.Pangkat) bool {
	query, err := config.DB.Exec(`UPDATE pangkats SET name = ?, updated_at = ? WHERE id = ?`, pangkat.Name, pangkat.UpdatedAt, id)
	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}
