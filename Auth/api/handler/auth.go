package handler

import (
	"auth/api/helper"
	"auth/api/helper/models"
	"auth/genproto/auth"
	"auth/token"
	"log"
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

	hashed_password, err := helper.HashPassword(request.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	code, err := h.SendEmail(request.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to send email. Please try again later.",
		})
		return
	}

	req := auth.Register{
		Id:          uuid.NewString(),
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		DateOfBirth: request.DateOfBirth,
		Gender:      request.Gender,
		Password:    hashed_password,
		Email:       request.Email,
		Role:        "user",
		Code:        code,
	}

	res, err := h.Auth.RegisterUser(ctx, &req)
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

	hashed_password, err := helper.HashPassword(request.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	code, err := h.SendEmail(request.Email)
	if err != nil {
		log.Println("----------------------------------->", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to send email. Please try again later.",
		})
		return
	}
	req := auth.Register{
		Id:          uuid.NewString(),
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		DateOfBirth: request.DateOfBirth,
		Gender:      request.Gender,
		Password:    hashed_password,
		Email:       request.Email,
		Role:        "courier",
		Code:        code,
	}

	res, err := h.Auth.RegisterUser(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// Login godoc
// @Summary Log in a courier
// @Description Authenticates a courier and returns a JWT token
// @Security BearerAuth
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body auth.Login true "Login credentials"
// @Success 200 {object} auth.TokenResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /auth/login [post]
func (h *Handler) Login(ctx *gin.Context) {
	var req auth.Login
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Auth.Login(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// GetUserProfile godoc
// @Summary Retrieves a user's profile
// @Description Fetches the user's profile information by their ID
// @Security BearerAuth
// @Tags Additional
// @Accept json
// @Produce json
// @Success 200 {object} auth.GetProfileResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /profile/get-profile [get]
func (h *Handler) GetProfile(ctx *gin.Context) {
	id, err := token.GetIdFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	req := auth.ById{Id: id}

	res, err := h.Auth.GetProfile(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// UpdateProfile godoc
// @Summary Updates a user's profile
// @Description Updates the user's profile information such as name, date of birth, and gender
// @Security BearerAuth
// @Tags Additional
// @Accept json
// @Produce json
// @Param body body models.UpdateRequest true "Additional update details"
// @Success 200 {object} auth.InfoResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /profile/update-profile [put]
func (h *Handler) UpdateProfile(ctx *gin.Context) {
	var req models.UpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := token.GetIdFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	request := auth.UpdateProfileRequest{
		Id:          id,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		DateOfBirth: req.DateOfBirth,
		Gender:      req.Gender,
	}

	res, err := h.Auth.UpdateProfile(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// RemoveProfile godoc
// @Summary Removes a user's profile
// @Description Deletes the user's profile based on their ID
// @Security BearerAuth
// @Tags Additional
// @Accept json
// @Produce json
// @Success 200 {object} auth.InfoResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /profile/delete-profile [delete]
func (h *Handler) DeleteProfile(ctx *gin.Context) {
	id, err := token.GetIdFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	req := auth.ById{Id: id}

	res, err := h.Auth.DeleteProfile(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
