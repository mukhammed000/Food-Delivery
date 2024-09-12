package postgres

import (
	"database/sql"
	"delivery/genproto/delivery"
	"fmt"
)

func (stg *DeliveryStorage) CreateOrder(req *delivery.CreateOrderRequest) (*delivery.InfoResponse, error) {
	query := `
        INSERT INTO orders (client_id, product_id, quantity, address, order_date, delivery_date, status)
        VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := stg.db.Exec(query, req.ClientId, req.ProductId, req.Quantity, req.Address, req.OrderDate, req.DeliveryDate, req.Status)
	if err != nil {
		return &delivery.InfoResponse{Success: false, Message: err.Error()}, err
	}

	return &delivery.InfoResponse{Success: true, Message: "Order created successfully"}, nil
}

func (stg *DeliveryStorage) GetOrder(req *delivery.GetOrderRequest) (*delivery.OrderResponse, error) {
	query := `
        SELECT order_id, client_id, product_id, quantity, address, order_date, delivery_date, status
        FROM orders
        WHERE order_id = $1`

	var order delivery.OrderResponse
	row := stg.db.QueryRow(query, req.OrderId)

	err := row.Scan(
		&order.OrderId,
		&order.ClientId,
		&order.ProductId,
		&order.Quantity,
		&order.Address,
		&order.OrderDate,
		&order.DeliveryDate,
		&order.Status,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("order not found")
	} else if err != nil {
		return nil, err
	}

	return &order, nil
}

func (stg *DeliveryStorage) GetAllOrders(req *delivery.GetAllOrdersRequest) (*delivery.GetAllOrdersResponse, error) {
	offset := (req.Page - 1) * req.Limit

	query := `
        SELECT order_id, client_id, product_id, quantity, address, order_date, delivery_date, status
        FROM orders
        ORDER BY order_date DESC
        LIMIT $1 OFFSET $2`

	rows, err := stg.db.Query(query, req.Limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*delivery.OrderResponse
	for rows.Next() {
		var order delivery.OrderResponse
		err := rows.Scan(
			&order.OrderId,
			&order.ClientId,
			&order.ProductId,
			&order.Quantity,
			&order.Address,
			&order.OrderDate,
			&order.DeliveryDate,
			&order.Status,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	countQuery := `SELECT COUNT(*) FROM orders`
	var totalCount int32
	err = stg.db.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, err
	}

	return &delivery.GetAllOrdersResponse{
		Orders:     orders,
		TotalCount: totalCount,
	}, nil
}

func (stg *DeliveryStorage) GetOrderByClient(req *delivery.GetOrderByClientRequest) (*delivery.GetOrderByClientResponse, error) {
	offset := (req.Page - 1) * req.Limit

	query := `
        SELECT order_id, client_id, product_id, quantity, address, order_date, delivery_date, status
        FROM orders
        WHERE client_id = $1
        ORDER BY order_date DESC
        LIMIT $2 OFFSET $3`

	rows, err := stg.db.Query(query, req.ClientId, req.Limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*delivery.OrderResponse
	for rows.Next() {
		var order delivery.OrderResponse
		err := rows.Scan(
			&order.OrderId,
			&order.ClientId,
			&order.ProductId,
			&order.Quantity,
			&order.Address,
			&order.OrderDate,
			&order.DeliveryDate,
			&order.Status,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	countQuery := `SELECT COUNT(*) FROM orders WHERE client_id = $1`
	var totalCount int32
	err = stg.db.QueryRow(countQuery, req.ClientId).Scan(&totalCount)
	if err != nil {
		return nil, err
	}

	return &delivery.GetOrderByClientResponse{
		Orders:     orders,
		TotalCount: totalCount,
	}, nil
}

func (stg *DeliveryStorage) UpdateOrder(req *delivery.UpdateOrderRequest) (*delivery.InfoResponse, error) {
	query := `
        UPDATE orders
        SET product_id = $1, quantity = $2, address = $3, delivery_date = $4, status = $5
        WHERE order_id = $6`

	_, err := stg.db.Exec(query, req.ProductId, req.Quantity, req.Address, req.DeliveryDate, req.Status, req.OrderId)
	if err != nil {
		return nil, err
	}

	return &delivery.InfoResponse{
		Message: "Order updated successfully",
	}, nil
}

func (stg *DeliveryStorage) DeleteOrder(req *delivery.DeleteOrderRequest) (*delivery.InfoResponse, error) {
	query := `
        DELETE FROM orders
        WHERE order_id = $1`

	result, err := stg.db.Exec(query, req.OrderId)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("no order found with id %s", req.OrderId)
	}

	return &delivery.InfoResponse{
		Message: "Order deleted successfully",
	}, nil
}
