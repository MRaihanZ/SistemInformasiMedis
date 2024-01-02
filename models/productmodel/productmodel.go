package productmodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func Getall() []entities.Product {
	rows, err := config.DB.Query(`
		SELECT 
			pasien.id, 
			pasien.name, 
			categories.name as category_name,
			pasien.created_at, 
			pasien.updated_at FROM pasien
		JOIN categories ON pasien.category_id = categories.id
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
			&product.Category.Name,
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
	result, err := config.DB.Exec(`
		INSERT INTO pasien(
			name, category_id, created_at, updated_at
		) VALUES (?, ?, ?, ?)`,
		product.Name,
		product.Category.Id,
		product.CreatedAt,
		product.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Detail(id int) entities.Product {
	row := config.DB.QueryRow(`
		SELECT 
			pasien.id, 
			pasien.name, 
			categories.name as category_name,
			pasien.created_at, 
			pasien.updated_at FROM pasien
		JOIN categories ON pasien.category_id = categories.id
		WHERE pasien.id = ?
	`, id)

	var product entities.Product

	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Category.Name,
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
		UPDATE pasien SET 
			name = ?, 
			category_id = ?,
			updated_at = ?
		WHERE id = ?`,
		product.Name,
		product.Category.Id,
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
	_, err := config.DB.Exec("DELETE FROM pasien WHERE id = ?", id)
	return err
}
