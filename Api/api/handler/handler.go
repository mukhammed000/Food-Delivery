package handler

import (
	pb "api/genproto/delivery"
)

type Handler struct {
	Cart    pb.CartServiceClient
	Office  pb.OfficeClient
	Order   pb.OrdersServiceClient
	Product pb.ProductServiceClient
}

func NewHandler(cart pb.CartServiceClient, Office pb.OfficeClient, Order pb.OrdersServiceClient, Product pb.ProductServiceClient) *Handler {
	return &Handler{
		Cart:    cart,
		Office:  Office,
		Order:   Order,
		Product: Product,
	}
}
