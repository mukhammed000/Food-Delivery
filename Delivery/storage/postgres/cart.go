package postgres

import (
	"database/sql"
	"delivery/genproto/delivery"
)

type DeliveryStorage struct {
	db *sql.DB
}

func NewDeliveryStorage(db *sql.DB) *DeliveryStorage {
	return &DeliveryStorage{
		db: db,
	}
}

func (stg *DeliveryStorage) CreateCart(req *delivery.CreateCartRequest) (*delivery.InfoResponse, error) {
	query := `INSERT INTO carts (id, user_id, status) VALUES ($1, $2, $3)`
	_, err := stg.db.Exec(query, req.Id, req.UserId, req.Status)
	if err != nil {
		return &delivery.InfoResponse{Success: false, Message: err.Error()}, err
	}
	return &delivery.InfoResponse{Success: true, Message: "Cart created successfully"}, nil
}

func (stg *DeliveryStorage) GetCart(req *delivery.ById) (*delivery.CartResponse, error) {
	query := `SELECT id, user_id, status FROM carts WHERE id = $1`
	row := stg.db.QueryRow(query, req.Id)

	var cart delivery.CartResponse
	var status string
	err := row.Scan(&cart.Id, &cart.UserId, &status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	cart.Status = status

	itemsQuery := `SELECT product_id, quantity, price FROM cart_items WHERE cart_id = $1`
	rows, err := stg.db.Query(itemsQuery, req.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*delivery.CartItem
	for rows.Next() {
		var item delivery.CartItem
		err := rows.Scan(&item.ProductId, &item.Quantity, &item.Price)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}
	cart.Items = items

	return &cart, nil
}

func (stg *DeliveryStorage) GetCartByUser(req *delivery.GetCartByUserRequest) (*delivery.CartResponse, error) {
	query := `SELECT id, user_id, status FROM carts WHERE user_id = $1 LIMIT 1`
	row := stg.db.QueryRow(query, req.UserId)

	var cart delivery.CartResponse
	var status string
	err := row.Scan(&cart.Id, &cart.UserId, &status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	cart.Status = status

	itemsQuery := `SELECT product_id, quantity, price FROM cart_items WHERE cart_id = $1`
	rows, err := stg.db.Query(itemsQuery, cart.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*delivery.CartItem
	for rows.Next() {
		var item delivery.CartItem
		err := rows.Scan(&item.ProductId, &item.Quantity, &item.Price)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}
	cart.Items = items

	return &cart, nil
}

func (stg *DeliveryStorage) UpdateCart(req *delivery.UpdateCartRequest) (*delivery.InfoResponse, error) {
	query := `UPDATE carts SET user_id = $1, status = $2 WHERE id = $3`
	result, err := stg.db.Exec(query, req.UserId, req.Status, req.Id)
	if err != nil {
		return &delivery.InfoResponse{Success: false, Message: err.Error()}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &delivery.InfoResponse{Success: false, Message: err.Error()}, err
	}

	if rowsAffected == 0 {
		return &delivery.InfoResponse{Success: false, Message: "Cart not found"}, nil
	}

	return &delivery.InfoResponse{Success: true, Message: "Cart updated successfully"}, nil
}

func (stg *DeliveryStorage) DeleteCart(req *delivery.ById) (*delivery.InfoResponse, error) {
	itemsQuery := `DELETE FROM cart_items WHERE cart_id = $1`
	_, err := stg.db.Exec(itemsQuery, req.Id)
	if err != nil {
		return &delivery.InfoResponse{Success: false, Message: err.Error()}, err
	}

	query := `DELETE FROM carts WHERE id = $1`
	result, err := stg.db.Exec(query, req.Id)
	if err != nil {
		return &delivery.InfoResponse{Success: false, Message: err.Error()}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &delivery.InfoResponse{Success: false, Message: err.Error()}, err
	}

	if rowsAffected == 0 {
		return &delivery.InfoResponse{Success: false, Message: "Cart not found"}, nil
	}

	return &delivery.InfoResponse{Success: true, Message: "Cart deleted successfully"}, nil
}

func (stg *DeliveryStorage) CreateCartItem(req *delivery.CreateCartItemRequest) (*delivery.InfoResponse, error) {
	query := `INSERT INTO cart_items (cart_id, product_id, quantity) VALUES ($1, $2, $3)`
	_, err := stg.db.Exec(query, req.CartId, req.ProductId, req.Quantity)
	if err != nil {
		return &delivery.InfoResponse{Success: false, Message: err.Error()}, err
	}
	return &delivery.InfoResponse{Success: true, Message: "Cart item created successfully"}, nil
}

func (stg *DeliveryStorage) GetCartItems(req *delivery.GetCartItemsRequest) (*delivery.GetCartItemsResponse, error) {
	query := `SELECT product_id, quantity, price FROM cart_items WHERE cart_id = $1`
	rows, err := stg.db.Query(query, req.CartId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*delivery.CartItem
	for rows.Next() {
		var item delivery.CartItem
		err := rows.Scan(&item.ProductId, &item.Quantity, &item.Price)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &delivery.GetCartItemsResponse{Items: items}, nil
}

func (stg *DeliveryStorage) DeleteCartItem(req *delivery.DeleteCartItemRequest) (*delivery.InfoResponse, error) {
	query := `DELETE FROM cart_items WHERE cart_id = $1 AND product_id = $2`
	result, err := stg.db.Exec(query, req.CartId, req.ProductId)
	if err != nil {
		return &delivery.InfoResponse{Success: false, Message: err.Error()}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &delivery.InfoResponse{Success: false, Message: err.Error()}, err
	}

	if rowsAffected == 0 {
		return &delivery.InfoResponse{Success: false, Message: "Cart item not found"}, nil
	}

	return &delivery.InfoResponse{Success: true, Message: "Cart item deleted successfully"}, nil
}

func (stg *DeliveryStorage) UpdateCartItemQuantity(req *delivery.UpdateCartItemQuantityRequest) (*delivery.InfoResponse, error) {
	query := `UPDATE cart_items SET quantity = $1 WHERE cart_id = $2 AND product_id = $3`
	result, err := stg.db.Exec(query, req.Quantity, req.CartId, req.ProductId)
	if err != nil {
		return &delivery.InfoResponse{Success: false, Message: err.Error()}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &delivery.InfoResponse{Success: false, Message: err.Error()}, err
	}

	if rowsAffected == 0 {
		return &delivery.InfoResponse{Success: false, Message: "Cart item not found"}, nil
	}

	return &delivery.InfoResponse{Success: true, Message: "Cart item quantity updated successfully"}, nil
}
