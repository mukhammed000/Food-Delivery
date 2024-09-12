package handler

import "github.com/gin-gonic/gin"

func (h *Handler) CreateOrder(ctx *gin.Context)      {}
func (h *Handler) GetOrder(ctx *gin.Context)         {}
func (h *Handler) GetAllOrders(ctx *gin.Context)     {}
func (h *Handler) GetOrderByClient(ctx *gin.Context) {}
func (h *Handler) UpdateOrder(ctx *gin.Context)      {}
func (h *Handler) DeleteOrder(ctx *gin.Context)      {}
