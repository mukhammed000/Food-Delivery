package postgres

import (
	"database/sql"
	"delivery/genproto/delivery"
	"fmt"
)

func (stg *DeliveryStorage) CreateProduct(req *delivery.CreateProductRequest) (*delivery.InfoResponse, error) {
	query := `
        INSERT INTO products (name, description, price, stock, category)
        VALUES ($1, $2, $3, $4, $5)`

	_, err := stg.db.Exec(query, req.Name, req.Description, req.Price, req.Stock, req.Category)
	if err != nil {
		return nil, err
	}

	return &delivery.InfoResponse{
		Message: "Product created successfully",
	}, nil
}

func (stg *DeliveryStorage) GetProduct(req *delivery.GetProductRequest) (*delivery.ProductResponse, error) {
	query := `SELECT product_id, name, description, price, stock, category FROM products WHERE product_id = $1`

	var product delivery.ProductResponse
	row := stg.db.QueryRow(query, req.ProductId)
	err := row.Scan(&product.ProductId, &product.Name, &product.Description, &product.Price, &product.Stock, &product.Category)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &product, nil
}

func (stg *DeliveryStorage) GetProducts(req *delivery.GetAllProductsRequest) (*delivery.GetProductsResponse, error) {
	offset := (req.Page - 1) * req.Limit

	query := `SELECT product_id, name, description, price, stock, category FROM products LIMIT $1 OFFSET $2`
	rows, err := stg.db.Query(query, req.Limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*delivery.ProductResponse

	for rows.Next() {
		var product delivery.ProductResponse
		if err := rows.Scan(&product.ProductId, &product.Name, &product.Description, &product.Price, &product.Stock, &product.Category); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	countQuery := `SELECT COUNT(*) FROM products`
	var totalCount int32
	err = stg.db.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, err
	}

	return &delivery.GetProductsResponse{
		Products:   products,
		TotalCount: totalCount,
	}, nil
}

func (stg *DeliveryStorage) UpdateProduct(req *delivery.UpdateProductRequest) (*delivery.InfoResponse, error) {
	if req.ProductId == "" {
		return nil, fmt.Errorf("ProductId is required")
	}

	query := `
        UPDATE products
        SET name = $1, description = $2, price = $3, stock = $4, category = $5
        WHERE product_id = $6
    `

	_, err := stg.db.Exec(query, req.Name, req.Description, req.Price, req.Stock, req.Category, req.ProductId)
	if err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}

	return &delivery.InfoResponse{
		Message: "Product updated successfully",
	}, nil
}

func (stg *DeliveryStorage) DeleteProduct(req *delivery.DeleteProductRequest) (*delivery.InfoResponse, error) {
	if req.ProductId == "" {
		return nil, fmt.Errorf("ProductId is required")
	}

	query := `
        DELETE FROM products
        WHERE product_id = $1
    `

	result, err := stg.db.Exec(query, req.ProductId)
	if err != nil {
		return nil, fmt.Errorf("failed to delete product: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("failed to check affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("product not found")
	}

	return &delivery.InfoResponse{
		Message: "Product deleted successfully",
	}, nil
}
