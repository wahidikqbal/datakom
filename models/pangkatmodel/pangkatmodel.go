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
