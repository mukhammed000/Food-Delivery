package handler

import (
	"api/genproto/delivery"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// CreateOrder handles the creation of a new order.
// @Summary Create a new order
// @Description Creates a new order based on the provided order details.
// @Security BearerAuth
// @Tags Orders
// @Accept json
// @Produce json
// @Param request body delivery.CreateOrderRequest true "Create order request"
// @Success 200 {object} delivery.InfoResponse "Order created successfully"
// @Failure 400 {object} delivery.InfoResponse "Invalid request"
// @Failure 500 {object} delivery.InfoResponse "Internal server error"
// @Router /order/create-order [post]
func (h *Handler) CreateOrder(ctx *gin.Context) {
	var req delivery.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, delivery.InfoResponse{Success: false, Message: "Invalid request"})
		return
	}

	resp, err := h.Order.CreateOrder(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, delivery.InfoResponse{Success: false, Message: "Failed to create order"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetOrder handles fetching details of a specific order by its ID.
// @Summary Get order details
// @Description Fetches details of an order based on the provided order ID.
// @Security BearerAuth
// @Tags Orders
// @Accept json
// @Produce json
// @Param order_id path string true "Order ID"
// @Success 200 {object} delivery.OrderResponse "Order details retrieved successfully"
// @Failure 400 {object} delivery.InfoResponse "Invalid request"
// @Failure 404 {object} delivery.InfoResponse "Order not found"
// @Failure 500 {object} delivery.InfoResponse "Internal server error"
// @Router /order/get-order [get]
func (h *Handler) GetOrder(ctx *gin.Context) {
	orderID := ctx.Param("order_id")
	if orderID == "" {
		ctx.JSON(http.StatusBadRequest, delivery.InfoResponse{Success: false, Message: "Order ID is required"})
		return
	}

	req := &delivery.GetOrderRequest{
		OrderId: orderID,
	}

	resp, err := h.Order.GetOrder(ctx, req)
	if err != nil {
		// Handle errors based on the gRPC error details
		if grpc.Code(err) == codes.NotFound {
			ctx.JSON(http.StatusNotFound, delivery.InfoResponse{Success: false, Message: "Order not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, delivery.InfoResponse{Success: false, Message: "Failed to retrieve order"})
		}
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetAllOrders handles fetching a paginated list of all orders.
// @Summary Get all orders
// @Description Retrieves a paginated list of all orders based on provided pagination details.
// @Security BearerAuth
// @Tags Orders
// @Accept json
// @Produce json
// @Param page query int true "Page number" default(1)
// @Param limit query int true "Number of orders per page" default(10)
// @Success 200 {object} delivery.GetAllOrdersResponse "Orders retrieved successfully"
// @Failure 400 {object} delivery.InfoResponse "Invalid request"
// @Failure 500 {object} delivery.InfoResponse "Internal server error"
// @Router /orders/get-all-orders [get]
func (h *Handler) GetAllOrders(ctx *gin.Context) {
	pageStr := ctx.Query("page")
	limitStr := ctx.Query("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		ctx.JSON(http.StatusBadRequest, delivery.InfoResponse{Success: false, Message: "Invalid page number"})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		ctx.JSON(http.StatusBadRequest, delivery.InfoResponse{Success: false, Message: "Invalid limit"})
		return
	}

	req := &delivery.GetAllOrdersRequest{
		Page:  int32(page),
		Limit: int32(limit),
	}

	resp, err := h.Order.GetAllOrders(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, delivery.InfoResponse{Success: false, Message: "Failed to retrieve orders"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetOrderByClient handles fetching orders by a specific client.
// @Summary Get orders by client
// @Description Retrieves a paginated list of orders for a specific client based on provided pagination details.
// @Security BearerAuth
// @Tags Orders
// @Accept json
// @Produce json
// @Param client_id query string true "Client ID"
// @Param page query int true "Page number" default(1)
// @Param limit query int true "Number of orders per page" default(10)
// @Success 200 {object} delivery.GetOrderByClientResponse "Orders retrieved successfully"
// @Failure 400 {object} delivery.InfoResponse "Invalid request"
// @Failure 500 {object} delivery.InfoResponse "Internal server error"
// @Router /orders/get-order-by-client [get]
func (h *Handler) GetOrderByClient(ctx *gin.Context) {
	clientID := ctx.Query("client_id")
	if clientID == "" {
		ctx.JSON(http.StatusBadRequest, delivery.InfoResponse{Success: false, Message: "Client ID is required"})
		return
	}

	pageStr := ctx.Query("page")
	limitStr := ctx.Query("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		ctx.JSON(http.StatusBadRequest, delivery.InfoResponse{Success: false, Message: "Invalid page number"})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		ctx.JSON(http.StatusBadRequest, delivery.InfoResponse{Success: false, Message: "Invalid limit"})
		return
	}

	req := &delivery.GetOrderByClientRequest{
		ClientId: clientID,
		Page:     int32(page),
		Limit:    int32(limit),
	}

	resp, err := h.Order.GetOrderByClient(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, delivery.InfoResponse{Success: false, Message: "Failed to retrieve orders"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// UpdateOrder handles updating an existing order.
// @Summary Update an existing order
// @Description Updates the details of an existing order based on the provided order ID and details.
// @Security BearerAuth
// @Tags Orders
// @Accept json
// @Produce json
// @Param request body delivery.UpdateOrderRequest true "Update order request"
// @Success 200 {object} delivery.InfoResponse "Order updated successfully"
// @Failure 400 {object} delivery.InfoResponse "Invalid request"
// @Failure 500 {object} delivery.InfoResponse "Internal server error"
// @Router /orders/update-order [put]
func (h *Handler) UpdateOrder(ctx *gin.Context) {
	var req delivery.UpdateOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, delivery.InfoResponse{Success: false, Message: "Invalid request"})
		return
	}

	resp, err := h.Order.UpdateOrder(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, delivery.InfoResponse{Success: false, Message: "Failed to update order"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// DeleteOrder handles the deletion of an order.
// @Summary Delete an order
// @Description Deletes an existing order based on the provided order ID.
// @Security BearerAuth
// @Tags Orders
// @Accept json
// @Produce json
// @Param request body delivery.DeleteOrderRequest true "Delete order request"
// @Success 200 {object} delivery.InfoResponse "Order deleted successfully"
// @Failure 400 {object} delivery.InfoResponse "Invalid request"
// @Failure 500 {object} delivery.InfoResponse "Internal server error"
// @Router /orders/delete-order [delete]
func (h *Handler) DeleteOrder(ctx *gin.Context) {
	var req delivery.DeleteOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, delivery.InfoResponse{Success: false, Message: "Invalid request"})
		return
	}

	resp, err := h.Order.DeleteOrder(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, delivery.InfoResponse{Success: false, Message: "Failed to delete order"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
