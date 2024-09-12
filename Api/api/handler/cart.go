package handler

import "github.com/gin-gonic/gin"

func (h *Handler) CreateCart(ctx *gin.Context)             {}
func (h *Handler) GetCart(ctx *gin.Context)                {}
func (h *Handler) GetCartByUser(ctx *gin.Context)          {}
func (h *Handler) UpdateCart(ctx *gin.Context)             {}
func (h *Handler) DeleteCart(ctx *gin.Context)             {}
func (h *Handler) CreateCartItem(ctx *gin.Context)         {}
func (h *Handler) GetCartItem(ctx *gin.Context)            {}
func (h *Handler) DeleteCartItem(ctx *gin.Context)         {}
func (h *Handler) UpdateCartItemQuantity(ctx *gin.Context) {}
