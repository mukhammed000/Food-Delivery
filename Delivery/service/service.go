package service

import (
	"context"
	"delivery/genproto/delivery"
	"delivery/storage"
	"log"
)

type DeliveryService struct {
	stg storage.InitRoot
	delivery.UnimplementedCartServiceServer
	delivery.UnimplementedOfficeServer
	delivery.UnimplementedOrdersServiceServer
	delivery.UnimplementedProductServiceServer
}

func NewDeliveryService(stg storage.InitRoot) *DeliveryService {
	return &DeliveryService{
		stg: stg,
	}
}

func (s *DeliveryService) CreateCart(ctx context.Context, req *delivery.CreateCartRequest) (*delivery.InfoResponse, error) {
	resp, err := s.stg.Delivery().CreateCart(req)
	if err != nil {
		log.Println("Error creating cart: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *DeliveryService) GetCart(ctx context.Context, req *delivery.ById) (*delivery.CartResponse, error) {
	resp, err := s.stg.Delivery().GetCart(req)
	if err != nil {
		log.Println("Error getting cart: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *DeliveryService) GetCartByUser(ctx context.Context, req *delivery.GetCartByUserRequest) (*delivery.CartResponse, error) {
	resp, err := s.stg.Delivery().GetCartByUser(req)
	if err != nil {
		log.Println("Error getting cart by user: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *DeliveryService) UpdateCart(ctx context.Context, req *delivery.UpdateCartRequest) (*delivery.InfoResponse, error) {
	resp, err := s.stg.Delivery().UpdateCart(req)
	if err != nil {
		log.Println("Error updating cart: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *DeliveryService) DeleteCart(ctx context.Context, req *delivery.ById) (*delivery.InfoResponse, error) {
	resp, err := s.stg.Delivery().DeleteCart(req)
	if err != nil {
		log.Println("Error deleting cart: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *DeliveryService) CreateCartItem(ctx context.Context, req *delivery.CreateCartItemRequest) (*delivery.InfoResponse, error) {
	resp, err := s.stg.Delivery().CreateCartItem(req)
	if err != nil {
		log.Println("Error creating cart item: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *DeliveryService) GetCartItems(ctx context.Context, req *delivery.GetCartItemsRequest) (*delivery.GetCartItemsResponse, error) {
	resp, err := s.stg.Delivery().GetCartItems(req)
	if err != nil {
		log.Println("Error getting cart items: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *DeliveryService) DeleteCartItem(ctx context.Context, req *delivery.DeleteCartItemRequest) (*delivery.InfoResponse, error) {
	resp, err := s.stg.Delivery().DeleteCartItem(req)
	if err != nil {
		log.Println("Error deleting cart item: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *DeliveryService) UpdateCartItemQuantity(ctx context.Context, req *delivery.UpdateCartItemQuantityRequest) (*delivery.InfoResponse, error) {
	resp, err := s.stg.Delivery().UpdateCartItemQuantity(req)
	if err != nil {
		log.Println("Error updating cart item quantity: ", err)
		return nil, err
	}
	return resp, nil
}

// Office Methods
func (s *DeliveryService) CreateOffice(ctx context.Context, req *delivery.CreateOfficeRequest) (*delivery.InfoResponse, error) {
	resp, err := s.stg.Delivery().CreateOffice(req)
	if err != nil {
		log.Println("Error creating office: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *DeliveryService) GetOffice(ctx context.Context, req *delivery.GetOfficeRequest) (*delivery.OfficeResponse, error) {
	resp, err := s.stg.Delivery().GetOffice(req)
	if err != nil {
		log.Println("Error getting office: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *DeliveryService) GetAllOffices(ctx context.Context, req *delivery.GetAllOfficesRequest) (*delivery.GetAllOfficesResponse, error) {
	resp, err := s.stg.Delivery().GetAllOffices(req)
	if err != nil {
		log.Println("Error getting all offices: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *DeliveryService) UpdateOffice(ctx context.Context, req *delivery.UpdateOfficeRequest) (*delivery.InfoResponse, error) {
	resp, err := s.stg.Delivery().UpdateOffice(req)
	if err != nil {
		log.Println("Error updating office: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *DeliveryService) DeleteOffice(ctx context.Context, req *delivery.DeleteOfficeRequest) (*delivery.InfoResponse, error) {
	resp, err := s.stg.Delivery().DeleteOffice(req)
	if err != nil {
		log.Println("Error deleting office: ", err)
		return nil, err
	}
	return resp, nil
}

// OrdersService Methods
func (s *DeliveryService) CreateOrder(ctx context.Context, req *delivery.CreateOrderRequest) (*delivery.InfoResponse, error) {
	resp, err := s.stg.Delivery().CreateOrder(req)
	if err != nil {
		log.Println("Error creating order: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *DeliveryService) GetOrder(ctx context.Context, req *delivery.GetOrderRequest) (*delivery.OrderResponse, error) {
	resp, err := s.stg.Delivery().GetOrder(req)
	if err != nil {
		log.Println("Error getting order: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *DeliveryService) GetAllOrders(ctx context.Context, req *delivery.GetAllOrdersRequest) (*delivery.GetAllOrdersResponse, error) {
	resp, err := s.stg.Delivery().GetAllOrders(req)
	if err != nil {
		log.Println("Error getting all orders: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *DeliveryService) GetOrderByClient(ctx context.Context, req *delivery.GetOrderByClientRequest) (*delivery.GetOrderByClientResponse, error) {
	resp, err := s.stg.Delivery().GetOrderByClient(req)
	if err != nil {
		log.Println("Error getting order by client: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *DeliveryService) UpdateOrder(ctx context.Context, req *delivery.UpdateOrderRequest) (*delivery.InfoResponse, error) {
	resp, err := s.stg.Delivery().UpdateOrder(req)
	if err != nil {
		log.Println("Error updating order: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *DeliveryService) DeleteOrder(ctx context.Context, req *delivery.DeleteOrderRequest) (*delivery.InfoResponse, error) {
	resp, err := s.stg.Delivery().DeleteOrder(req)
	if err != nil {
		log.Println("Error deleting order: ", err)
		return nil, err
	}
	return resp, nil
}

// ProductService Methods
func (s *DeliveryService) CreateProduct(ctx context.Context, req *delivery.CreateProductRequest) (*delivery.InfoResponse, error) {
	resp, err := s.stg.Delivery().CreateProduct(req)
	if err != nil {
		log.Println("Error creating product: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *DeliveryService) GetProduct(ctx context.Context, req *delivery.GetProductRequest) (*delivery.ProductResponse, error) {
	resp, err := s.stg.Delivery().GetProduct(req)
	if err != nil {
		log.Println("Error getting product: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *DeliveryService) GetProducts(ctx context.Context, req *delivery.GetAllProductsRequest) (*delivery.GetProductsResponse, error) {
	resp, err := s.stg.Delivery().GetProducts(req)
	if err != nil {
		log.Println("Error getting all products: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *DeliveryService) UpdateProduct(ctx context.Context, req *delivery.UpdateProductRequest) (*delivery.InfoResponse, error) {
	resp, err := s.stg.Delivery().UpdateProduct(req)
	if err != nil {
		log.Println("Error updating product: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *DeliveryService) DeleteProduct(ctx context.Context, req *delivery.DeleteProductRequest) (*delivery.InfoResponse, error) {
	resp, err := s.stg.Delivery().DeleteProduct(req)
	if err != nil {
		log.Println("Error deleting product: ", err)
		return nil, err
	}
	return resp, nil
}
