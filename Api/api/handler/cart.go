package handler

import (
	"api/genproto/delivery"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCart handles the creation of a new cart.
// @Summary Create a new cart
// @Description Creates a new cart for a user and returns the cart details.
// @Tags Cart
// @Accept json
// @Produce json
// @Param request body delivery.CreateCartRequest true "Cart creation request"
// @Success 200 {object} delivery.CartResponse "Cart created successfully"
// @Failure 400 {object} delivery.InfoResponse"Bad request"
// @Failure 500 {object} delivery.InfoResponse"Internal server error"
// @Router /cart/create-cart [post]
func (h *Handler) CreateCart(ctx *gin.Context) {
	var req delivery.CreateCartRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, delivery.InfoResponse{Message: "Invalid request", Success: false})
		return
	}

	resp, err := h.Cart.CreateCart(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, delivery.InfoResponse{Message: "Failed to create cart"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetCart handles retrieving a cart by ID.
// @Summary Get a cart by ID
// @Description Retrieves the details of a cart based on the provided cart ID.
// @Tags Cart
// @Accept json
// @Produce json
// @Param id path string true "Cart ID"
// @Success 200 {object} delivery.CartResponse "Cart retrieved successfully"
// @Failure 400 {object} delivery.InfoResponse "Bad request"
// @Failure 404 {object} delivery.InfoResponse "Cart not found"
// @Failure 500 {object} delivery.InfoResponse "Internal server error"
// @Router /cart/{id} [get]
func (h *Handler) GetCart(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.Cart.GetCart(ctx, &delivery.ById{Id: id})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, delivery.InfoResponse{Message: "Failed to retrieve cart"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetCartByUser handles retrieving a cart by user ID.
// @Summary Get a cart by user ID
// @Description Retrieves the details of a cart for a specific user based on the provided user ID.
// @Tags Cart
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} delivery.CartResponse "Cart retrieved successfully"
// @Failure 400 {object} delivery.InfoResponse "Bad request"
// @Failure 404 {object} delivery.InfoResponse "Cart not found"
// @Failure 500 {object} delivery.InfoResponse "Internal server error"
// @Router /cart/user/{user_id} [get]
func (h *Handler) GetCartByUser(ctx *gin.Context) {
	userID := ctx.Param("user_id")

	resp, err := h.Cart.GetCartByUser(ctx, &delivery.GetCartByUserRequest{UserId: userID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, delivery.InfoResponse{Message: "Failed to retrieve cart"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// UpdateCart handles updating an existing cart.
// @Summary Update an existing cart
// @Description Updates the details of an existing cart based on the provided cart ID and request data.
// @Tags Cart
// @Accept json
// @Produce json
// @Param id path string true "Cart ID"
// @Param request body delivery.UpdateCartRequest true "Cart update request"
// @Success 200 {object} delivery.InfoResponse "Cart updated successfully"
// @Failure 400 {object} delivery.InfoResponse "Bad request"
// @Failure 404 {object} delivery.InfoResponse "Cart not found"
// @Failure 500 {object} delivery.InfoResponse "Internal server error"
// @Router /cart/{id} [put]
func (h *Handler) UpdateCart(ctx *gin.Context) {
	id := ctx.Param("id")

	var req delivery.UpdateCartRequest
	req.Id = id

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, delivery.InfoResponse{Message: "Invalid request", Success: false})
		return
	}

	resp, err := h.Cart.UpdateCart(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, delivery.InfoResponse{Message: "Failed to update cart"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// DeleteCart handles deleting a cart by ID.
// @Summary Delete a cart by ID
// @Description Deletes an existing cart based on the provided cart ID.
// @Tags Cart
// @Accept json
// @Produce json
// @Param id path string true "Cart ID"
// @Success 200 {object} delivery.InfoResponse "Cart deleted successfully"
// @Failure 400 {object} delivery.InfoResponse "Bad request"
// @Failure 500 {object} delivery.InfoResponse "Internal server error"
// @Router /cart/{id} [delete]
func (h *Handler) DeleteCart(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.Cart.DeleteCart(ctx, &delivery.ById{Id: id})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, delivery.InfoResponse{Message: "Failed to delete cart", Success: false})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// CreateCartItem handles adding a new item to a cart.
// @Summary Add a new item to a cart
// @Description Adds an item to the specified cart based on the provided details.
// @Tags Cart
// @Accept json
// @Produce json
// @Param request body delivery.CreateCartItemRequest true "Cart item creation request"
// @Success 200 {object} delivery.InfoResponse "Item added to cart successfully"
// @Failure 400 {object} delivery.InfoResponse "Bad request"
// @Failure 500 {object} delivery.InfoResponse "Internal server error"
// @Router /cart/item [post]
func (h *Handler) CreateCartItem(ctx *gin.Context) {
	var req delivery.CreateCartItemRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, delivery.InfoResponse{Message: "Invalid request", Success: false})
		return
	}

	resp, err := h.Cart.CreateCartItem(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, delivery.InfoResponse{Message: "Failed to add item to cart", Success: false})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetCartItem handles retrieving a cart item based on cart ID and product ID.
// @Summary Get a cart item by cart ID and product ID
// @Description Retrieves the details of a specific item in a cart based on the provided cart ID and product ID.
// @Tags Cart
// @Accept json
// @Produce json
// @Param cart_id path string true "Cart ID"
// @Param product_id path string true "Product ID"
// @Success 200 {object} delivery.CartItem "Cart item details"
// @Failure 400 {object} delivery.InfoResponse "Bad request"
// @Failure 404 {object} delivery.InfoResponse "Cart item not found"
// @Failure 500 {object} delivery.InfoResponse "Internal server error"
// @Router /cart/item/{cart_id}/{product_id} [get]
func (h *Handler) GetCartItem(ctx *gin.Context) {
	cartID := ctx.Param("cart_id")
	productID := ctx.Param("product_id")

	resp, err := h.Cart.GetCartItems(ctx, &delivery.GetCartItemsRequest{CartId: cartID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, delivery.InfoResponse{Message: "Failed to retrieve cart item", Success: false})
		return
	}

	var item *delivery.CartItem
	for _, cartItem := range resp.Items {
		if cartItem.ProductId == productID {
			item = cartItem
			break
		}
	}

	if item != nil {
		ctx.JSON(http.StatusOK, item)
	} else {
		ctx.JSON(http.StatusNotFound, resp)
	}
}

// DeleteCartItem handles the deletion of a specific item from a cart.
// @Summary Delete a cart item
// @Description Deletes a specific item from a cart based on the provided cart ID and product ID.
// @Tags Cart
// @Accept json
// @Produce json
// @Param cart_id path string true "Cart ID"
// @Param product_id path string true "Product ID"
// @Success 200 {object} delivery.InfoResponse "Cart item deleted successfully"
// @Failure 400 {object} delivery.InfoResponse "Bad request"
// @Failure 404 {object} delivery.InfoResponse "Cart item not found"
// @Failure 500 {object} delivery.InfoResponse "Internal server error"
// @Router /cart/item/{cart_id}/{product_id} [delete]
func (h *Handler) DeleteCartItem(ctx *gin.Context) {
	cartID := ctx.Param("cart_id")
	productID := ctx.Param("product_id")

	resp, err := h.Cart.DeleteCartItem(ctx, &delivery.DeleteCartItemRequest{
		CartId:    cartID,
		ProductId: productID,
	})
	if err != nil {
		if err.Error() == "cart item not found" {
			ctx.JSON(http.StatusNotFound, delivery.InfoResponse{Message: "Cart item not found", Success: false})
		} else {
			ctx.JSON(http.StatusInternalServerError, delivery.InfoResponse{Message: "Failed to delete cart item", Success: false})
		}
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// UpdateCartItemQuantity handles the update of the quantity of a specific item in a cart.
// @Summary Update the quantity of a cart item
// @Description Updates the quantity of a specific item in a cart based on the provided cart ID, product ID, and new quantity.
// @Tags Cart
// @Accept json
// @Produce json
// @Param cart_id path string true "Cart ID"
// @Param product_id path string true "Product ID"
// @Param request body delivery.UpdateCartItemQuantityRequest true "Update cart item quantity request"
// @Success 200 {object} delivery.InfoResponse "Cart item quantity updated successfully"
// @Failure 400 {object} delivery.InfoResponse "Bad request"
// @Failure 404 {object} delivery.InfoResponse "Cart item not found"
// @Failure 500 {object} delivery.InfoResponse "Internal server error"
// @Router /cart/item/{cart_id}/{product_id}/quantity [put]
func (h *Handler) UpdateCartItemQuantity(ctx *gin.Context) {
	cartID := ctx.Param("cart_id")
	productID := ctx.Param("product_id")

	var req delivery.UpdateCartItemQuantityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, delivery.InfoResponse{Message: "Invalid request", Success: false})
		return
	}

	req.CartId = cartID
	req.ProductId = productID

	resp, err := h.Cart.UpdateCartItemQuantity(ctx, &req)
	if err != nil {
		if err.Error() == "cart item not found" {
			ctx.JSON(http.StatusNotFound, delivery.InfoResponse{Message: "Cart item not found", Success: false})
		} else {
			ctx.JSON(http.StatusInternalServerError, delivery.InfoResponse{Message: "Failed to update cart item quantity", Success: false})
		}
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
