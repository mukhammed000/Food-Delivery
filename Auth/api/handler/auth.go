package handler

import (
	"auth/api/helper/models"
	"auth/genproto/auth"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user with email, password, and personal details
// @Security BearerAuth
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body models.UserRegister true "User registration details"
// @Success 200 {object} auth.InfoResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /auth/register-user [post]
func (h *Handler) RegisterUser(ctx *gin.Context) {
	var request models.UserRegister
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req := auth.Register{
		Id:          uuid.NewString(),
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		DateOfBirth: request.DateOfBirth,
		Gender:      request.Gender,
		Password:    request.Password,
		Email:       request.Email,
		Role:        "user",
	}

	res, err := h.Auth.RegisterUser(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// LoginUser godoc
// @Summary Log in a user
// @Description Authenticates a user and returns a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body auth.Login true "User login credentials"
// @Success 200 {object} auth.TokenResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /auth/login-user [post]
func (h *Handler) LoginUser(ctx *gin.Context) {
	var req auth.Login
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Auth.LoginUser(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// RegisterCourier godoc
// @Summary Register a new courier
// @Description Register a new courier with email, password, and personal details
// @Security BearerAuth
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body models.UserRegister true "Courier registration details"
// @Success 200 {object} auth.InfoResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /auth/register-courier [post]
func (h *Handler) RegisterCourier(ctx *gin.Context) {
	var request models.UserRegister
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req := auth.Register{
		Id:          uuid.NewString(),
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		DateOfBirth: request.DateOfBirth,
		Gender:      request.Gender,
		Password:    request.Password,
		Email:       request.Email,
		Role:        "courier",
	}

	res, err := h.Auth.RegisterUser(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// LoginUser godoc
// @Summary Log in a courier
// @Description Authenticates a courier and returns a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body auth.Login true "User login credentials"
// @Success 200 {object} auth.TokenResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /auth/login-courier [post]
func (h *Handler) LoginCourier(ctx *gin.Context) {
	var req auth.Login
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Auth.LoginUser(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
