package stockmodel

import (
	"go-web/config"
	"go-web/entities"
)

func GetAll() []entities.Stock {
	rows, err := config.DB.Query(`SELECT * FROM stocks`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var stocks []entities.Stock

	for rows.Next() {
		var stock entities.Stock
		if err := rows.Scan(&stock.Id, &stock.Name, &stock.CreatedAt, &stock.CreatedAt); err != nil {
			panic(err)
		}

		stocks = append(stocks, stock)
	}
	return stocks
}

func Create(stock entities.Stock) bool {
	result, err := config.DB.Exec(`INSERT INTO stocks (name, created_at, updated_at) VALUE (?, ?, ?)`, stock.Name, stock.UpdatedAt, stock.CreatedAt)
	if err != nil {
		panic(err)
	}

	isInsert, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return isInsert > 0

}

func Detail(id int) entities.Stock {
	rows := config.DB.QueryRow(`SELECT id, name FROM stocks WHERE id = ?`, id)

	var stock entities.Stock
	if err := rows.Scan(&stock.Id, &stock.Name); err != nil {
		panic(err)
	}

	return stock
}

func Update(id int, stock entities.Stock) bool {
	result, err := config.DB.Exec(`UPDATE stocks SET name = ?, updated_at = ? WHERE id = ?`, stock.Name, stock.UpdatedAt, id)
	if err != nil {
		panic(err)
	}

	isUpdate, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	return isUpdate > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec(`DELETE FROM stocks WHERE id = ?`, id)
	return err
}
