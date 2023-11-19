package productmodel

import (
	"go-web/config"
	"go-web/entities"
)

func GetAll() []entities.Product {
	rows, err := config.DB.Query(`
		SELECT 
			products.id, 
			products.name,
			pangkats.name as pangkat_name,
			products.nrp,
			units.name as unit_name,
			categories.name as category_name,
			products.serialnumber,
			stocks.name as stock_name, 
			products.created_at, 
			products.updated_at
		FROM products
		JOIN pangkats ON products.pangkat_id = pangkats.id
		JOIN units ON products.unit_id = units.id
		JOIN categories ON products.category_id = categories.id
		JOIN stocks ON products.stock_id = stocks.id
	`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var products []entities.Product

	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Pangkat.Name,
			&product.Nrp,
			&product.Unit.Name,
			&product.Category.Name,
			&product.Serialnumber,
			&product.Stock.Name,
			&product.CreatedAt,
			&product.UpdatedAt,
		); err != nil {
			panic(err)
		}

		products = append(products, product)
	}

	return products
}
