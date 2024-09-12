package storage

import "delivery/genproto/delivery"

type InitRoot interface {
	Delivery() DeliveryService
}

type DeliveryService interface {
	// CartService Methods
	CreateCart(req *delivery.CreateCartRequest) (*delivery.InfoResponse, error)
	GetCart(req *delivery.ById) (*delivery.CartResponse, error)
	GetCartByUser(req *delivery.GetCartByUserRequest) (*delivery.CartResponse, error)
	UpdateCart(req *delivery.UpdateCartRequest) (*delivery.InfoResponse, error)
	DeleteCart(req *delivery.ById) (*delivery.InfoResponse, error)
	CreateCartItem(req *delivery.CreateCartItemRequest) (*delivery.InfoResponse, error)
	GetCartItems(req *delivery.GetCartItemsRequest) (*delivery.GetCartItemsResponse, error)
	DeleteCartItem(req *delivery.DeleteCartItemRequest) (*delivery.InfoResponse, error)
	UpdateCartItemQuantity(req *delivery.UpdateCartItemQuantityRequest) (*delivery.InfoResponse, error)

	// Office Methods
	CreateOffice(req *delivery.CreateOfficeRequest) (*delivery.InfoResponse, error)
	GetOffice(req *delivery.GetOfficeRequest) (*delivery.OfficeResponse, error)
	GetAllOffices(req *delivery.GetAllOfficesRequest) (*delivery.GetAllOfficesResponse, error)
	UpdateOffice(req *delivery.UpdateOfficeRequest) (*delivery.InfoResponse, error)
	DeleteOffice(req *delivery.DeleteOfficeRequest) (*delivery.InfoResponse, error)

	// OrdersService Methods
	CreateOrder(req *delivery.CreateOrderRequest) (*delivery.InfoResponse, error)
	GetOrder(req *delivery.GetOrderRequest) (*delivery.OrderResponse, error)
	GetAllOrders(req *delivery.GetAllOrdersRequest) (*delivery.GetAllOrdersResponse, error)
	GetOrderByClient(req *delivery.GetOrderByClientRequest) (*delivery.GetOrderByClientResponse, error)
	UpdateOrder(req *delivery.UpdateOrderRequest) (*delivery.InfoResponse, error)
	DeleteOrder(req *delivery.DeleteOrderRequest) (*delivery.InfoResponse, error)

	// ProductService Methods
	CreateProduct(req *delivery.CreateProductRequest) (*delivery.InfoResponse, error)
	GetProduct(req *delivery.GetProductRequest) (*delivery.ProductResponse, error)
	GetProducts(req *delivery.GetAllProductsRequest) (*delivery.GetProductsResponse, error)
	UpdateProduct(req *delivery.UpdateProductRequest) (*delivery.InfoResponse, error)
	DeleteProduct(req *delivery.DeleteProductRequest) (*delivery.InfoResponse, error)
}
