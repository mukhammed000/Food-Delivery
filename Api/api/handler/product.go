package handler

import (
	"api/genproto/delivery"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateProduct handles the creation of a new product.
// @Summary Create a new product
// @Description Creates a new product based on the provided product details.
// @Security BearerAuth
// @Tags Product
// @Accept json
// @Produce json
// @Param request body delivery.CreateProductRequest true "Create product request"
// @Success 200 {object} delivery.InfoResponse "Product created successfully"
// @Failure 400 {object} delivery.InfoResponse "Invalid request"
// @Failure 500 {object} delivery.InfoResponse "Internal server error"
// @Router /product/create-product [post]
func (h *Handler) CreateProduct(ctx *gin.Context) {
	var req delivery.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, delivery.InfoResponse{Success: false, Message: "Invalid request"})
		return
	}

	resp, err := h.Product.CreateProduct(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, delivery.InfoResponse{Success: false, Message: "Failed to create product"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetProduct handles retrieving a product by its ID.
// @Summary Get a product by ID
// @Description Retrieves a product based on the provided product ID.
// @Security BearerAuth
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} delivery.ProductResponse "Product retrieved successfully"
// @Failure 400 {object} delivery.InfoResponse "Invalid request"
// @Failure 404 {object} delivery.InfoResponse "Product not found"
// @Failure 500 {object} delivery.InfoResponse "Internal server error"
// @Router /product/get-product [get]
func (h *Handler) GetProduct(ctx *gin.Context) {
	productID := ctx.Param("id")

	if productID == "" {
		ctx.JSON(http.StatusBadRequest, delivery.InfoResponse{Success: false, Message: "Product ID is required"})
		return
	}

	req := &delivery.GetProductRequest{
		ProductId: productID,
	}

	resp, err := h.Product.GetProduct(ctx, req)
	if err != nil {
		if err.Error() == "product not found" {
			ctx.JSON(http.StatusNotFound, delivery.InfoResponse{Success: false, Message: "Product not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, delivery.InfoResponse{Success: false, Message: "Failed to retrieve product"})
		}
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetAllProduct handles retrieving a list of all products with optional pagination.
// @Summary Get all products
// @Description Retrieves a list of all products with optional pagination.
// @Security BearerAuth
// @Tags Product
// @Accept json
// @Produce json
// @Param page query int32 false "Page number" default(1)
// @Param limit query int32 false "Number of products per page" default(10)
// @Success 200 {object} delivery.GetProductsResponse "Products retrieved successfully"
// @Failure 400 {object} delivery.InfoResponse "Invalid request"
// @Failure 500 {object} delivery.InfoResponse "Internal server error"
// @Router /product/get-all-products [get]
func (h *Handler) GetAllProduct(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	req := &delivery.GetAllProductsRequest{
		Page:  int32(page),
		Limit: int32(limit),
	}

	resp, err := h.Product.GetProducts(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, delivery.InfoResponse{Success: false, Message: "Failed to retrieve products"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// UpdateProduct handles updating an existing product.
// @Summary Update an existing product
// @Description Updates an existing product based on the provided product details.
// @Security BearerAuth
// @Tags Product
// @Accept json
// @Produce json
// @Param request body delivery.UpdateProductRequest true "Update product request"
// @Success 200 {object} delivery.InfoResponse "Product updated successfully"
// @Failure 400 {object} delivery.InfoResponse "Invalid request"
// @Failure 500 {object} delivery.InfoResponse "Internal server error"
// @Router /product/update-product [put]
func (h *Handler) UpdateProduct(ctx *gin.Context) {
	var req delivery.UpdateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, delivery.InfoResponse{Success: false, Message: "Invalid request"})
		return
	}

	resp, err := h.Product.UpdateProduct(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, delivery.InfoResponse{Success: false, Message: "Failed to update product"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// DeleteProduct handles the deletion of a product.
// @Summary Delete a product
// @Description Deletes a product based on the provided product ID.
// @Security BearerAuth
// @Tags Product
// @Accept json
// @Produce json
// @Param request body delivery.DeleteProductRequest true "Delete product request"
// @Success 200 {object} delivery.InfoResponse "Product deleted successfully"
// @Failure 400 {object} delivery.InfoResponse "Invalid request"
// @Failure 500 {object} delivery.InfoResponse "Internal server error"
// @Router /product/delete-product [delete]
func (h *Handler) DeleteProduct(ctx *gin.Context) {
	var req delivery.DeleteProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, delivery.InfoResponse{Success: false, Message: "Invalid request"})
		return
	}

	resp, err := h.Product.DeleteProduct(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, delivery.InfoResponse{Success: false, Message: "Failed to delete product"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
