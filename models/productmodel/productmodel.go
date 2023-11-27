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
			products.created_at, 
			products.updated_at
		FROM products
		JOIN pangkats ON products.pangkat_id = pangkats.id
		JOIN units ON products.unit_id = units.id
		JOIN categories ON products.category_id = categories.id
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
			&product.CreatedAt,
			&product.UpdatedAt,
		); err != nil {
			panic(err)
		}

		products = append(products, product)
	}

	return products
}

func Create(product entities.Product) bool {
	result, err := config.DB.Exec(`INSERT INTO products (
		name, pangkat_id, nrp, unit_id, category_id, serialnumber, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		product.Name,
		product.Pangkat.Id,
		product.Nrp,
		product.Unit.Id,
		product.Category.Id,
		product.Serialnumber,
		product.CreatedAt,
		product.UpdatedAt,
	)
	if err != nil {
		panic(err)
	}

	oke, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return oke > 0
}

func Detail(id int) entities.Product {
	row := config.DB.QueryRow(`
	SELECT 
	products.id, 
	products.name,
	pangkats.name as pangkat_name,
	products.nrp,
	units.name as unit_name,
	categories.name as category_name,
	products.serialnumber,
	products.created_at, 
	products.updated_at
	FROM products
	JOIN pangkats ON products.pangkat_id = pangkats.id
	JOIN units ON products.unit_id = units.id
	JOIN categories ON products.category_id = categories.id
	WHERE products.id = ?
	`, id)

	var product entities.Product
	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Pangkat.Name,
		&product.Nrp,
		&product.Unit.Name,
		&product.Category.Name,
		&product.Serialnumber,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	return product
}

func Update(id int, product entities.Product) bool {
	query, err := config.DB.Exec(`
		UPDATE products SET
			name = ?,
			pangkat_id = ?,
			nrp = ?,
			unit_id = ?,
			category_id = ?,
			serialnumber = ?,
			updated_at = ?
		WHERE id = ?
		`,
		product.Name,
		product.Pangkat.Id,
		product.Nrp,
		product.Unit.Id,
		product.Category.Id,
		product.Serialnumber,
		product.UpdatedAt,
		id,
	)

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
	_, err := config.DB.Exec(`DELETE FROM products WHERE id = ?`, id)
	return err
}
