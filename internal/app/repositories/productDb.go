package repositories

import (
	"time"

	"github.com/raexera/soko/internal/app/models"
	"github.com/raexera/soko/internal/app/storage"
)

func CreateProduct(product models.Product) (models.Product, error) {
	db := storage.GetDB()

	sqlStatement := `INSERT INTO products (name, description, category, quantity, price) VALUES ($1, $2, $3, $4, $5)  RETURNING id`

	err := db.QueryRow(
		sqlStatement,
		product.Name,
		product.Description,
		product.Category,
		product.Quantity,
		product.Price,
	).Scan(&product.Id)
	if err != nil {
		return product, err
	}
	return product, nil
}

func GetProductByID(id int) (models.Product, error) {
	db := storage.GetDB()

	var product models.Product

	sqlStatement := `SELECT id, name, description, category, quantity, price FROM products WHERE id = $1`

	err := db.QueryRow(sqlStatement, id).
		Scan(&product.Id, &product.Name, &product.Description, &product.Category, &product.Quantity, &product.Price)
	if err != nil {
		return product, err
	}

	return product, nil
}

func GetAllProducts() ([]models.Product, error) {
	db := storage.GetDB()
	var products []models.Product

	sqlStatement := `SELECT id, name, description, category, quantity, price FROM products`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Category, &product.Quantity, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func UpdateProduct(product models.Product, id int) (models.Product, error) {
	db := storage.GetDB()
	sqlStatement := `
    UPDATE products
    SET name = $2, description = $3, category = $4, quantity = $5, price = $6, updated_at = $7
    WHERE id = $1
    RETURNING id`
	err := db.QueryRow(sqlStatement, id, product.Name, product.Description, product.Category, product.Quantity, product.Price, time.Now()).
		Scan(&product.Id)
	if err != nil {
		return models.Product{}, err
	}
	product.Id = id
	return product, nil
}

func DeleteProduct(id int) error {
	db := storage.GetDB()
	sqlStatement := `DELETE FROM products WHERE id = $1`
	_, err := db.Exec(sqlStatement, id)
	return err
}
