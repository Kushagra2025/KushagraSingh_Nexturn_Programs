package repository

import (
	"E_Commerce_Application/model"
	"database/sql"
	"log"
)

type ProductRepository struct {
	db *sql.DB
}

// NewProductRepository creates a new instance of ProductRepository.
func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db}
}

// CreateProduct adds a new product to the database.
func (r *ProductRepository) CreateProduct(product *model.Product) error {
	query := "INSERT INTO products (name, description, price, stock, category_id) VALUES (?, ?, ?, ?, ?)"
	_, err := r.db.Exec(query, product.Name, product.Description, product.Price, product.Stock, product.CategoryID)
	if err != nil {
		log.Println("Error inserting product: ", err)
		return err
	}
	return nil
}

// GetProduct fetches a product by ID.
func (r *ProductRepository) GetProduct(id int) (*model.Product, error) {
	query := "SELECT id, name, description, price, stock, category_id FROM products WHERE id = ?"
	row := r.db.QueryRow(query, id)
	var product model.Product
	if err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CategoryID); err != nil {
		log.Println("Error fetching product: ", err)
		return nil, err
	}
	return &product, nil
}

// UpdateProductStock updates the stock of a product.
func (r *ProductRepository) UpdateProductStock(id int, stock int) error {
	query := "UPDATE products SET stock = ? WHERE id = ?"
	_, err := r.db.Exec(query, stock, id)
	if err != nil {
		log.Println("Error updating stock: ", err)
		return err
	}
	return nil
}

// DeleteProduct deletes a product by ID.
func (r *ProductRepository) DeleteProduct(id int) error {
	query := "DELETE FROM products WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Println("Error deleting product: ", err)
		return err
	}
	return nil
}
