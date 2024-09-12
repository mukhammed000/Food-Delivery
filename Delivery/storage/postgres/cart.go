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
	return nil, nil
}

func (stg *DeliveryStorage) GetCart(req *delivery.ById) (*delivery.CartResponse, error) {
	return nil, nil
}

func (stg *DeliveryStorage) GetCartByUser(req *delivery.GetCartByUserRequest) (*delivery.CartResponse, error) {
	return nil, nil
}

func (stg *DeliveryStorage) UpdateCart(req *delivery.UpdateCartRequest) (*delivery.InfoResponse, error) {
	return nil, nil
}

func (stg *DeliveryStorage) DeleteCart(req *delivery.ById) (*delivery.InfoResponse, error) {
	return nil, nil
}

func (stg *DeliveryStorage) CreateCartItem(req *delivery.CreateCartItemRequest) (*delivery.InfoResponse, error) {
	return nil, nil
}

func (stg *DeliveryStorage) GetCartItems(req *delivery.GetCartItemsRequest) (*delivery.GetCartItemsResponse, error) {
	return nil, nil
}

func (stg *DeliveryStorage) DeleteCartItem(req *delivery.DeleteCartItemRequest) (*delivery.InfoResponse, error) {
	return nil, nil
}

func (stg *DeliveryStorage) UpdateCartItemQuantity(req *delivery.UpdateCartItemQuantityRequest) (*delivery.InfoResponse, error) {
	return nil, nil
}
