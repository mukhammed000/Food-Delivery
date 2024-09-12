package handler

import (
	"auth/genproto/auth"
	"auth/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ChangeUserPassword godoc
// @Summary Changes a user's password
// @Description Updates the user's password after validating the current password
// @Security BearerAuth
// @Tags Additional
// @Accept json
// @Produce json
// @Param current_password query string true "Current Password"
// @Param new_password query string true "New Password"
// @Success 200 {object} auth.InfoResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /additional/change-password [put]
func (h *Handler) ChangePassword(ctx *gin.Context) {
	current_password := ctx.Query("current_password")
	new_password := ctx.Query("new_password")

	id, err := token.GetIdFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	req := auth.ChangePasswordRequest{
		CurrentPassword: current_password,
		NewPassword:     new_password,
		Id:              id,
	}

	res, err := h.Auth.ChangePassword(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// ConfirmEmail godoc
// @Summary Confirms a user's email address
// @Description Verifies the user's email address using a verification code
// @Security BearerAuth
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body auth.VerifyEmailRequest true "Email verification details"
// @Success 200 {object} auth.TokenResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /auth/verify-email [post]
func (h *Handler) VerifyEmail(ctx *gin.Context) {
	var req auth.VerifyEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Auth.VerifyEmail(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// UpdateEmail godoc
// @Summary Updates a user's email address
// @Description Changes the user's email address and sends a verification code
// @Security BearerAuth
// @Tags Additional
// @Accept json
// @Produce json
// @Param new_email query string true "New Email"
// @Param password query string true "Password"
// @Success 200 {object} auth.InfoResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /additional/change-email [put]
func (h *Handler) ChangeEmail(ctx *gin.Context) {
	var check auth.CheckEmailAndPasswordRequest
	new_email := ctx.Query("new_email")
	password := ctx.Query("password")

	current_email, err := token.GetEmailFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	req := auth.ChangeEmailRequest{
		CurrentEmail: current_email,
		NewEmail:     new_email,
		Password:     password,
	}

	check.Email = req.CurrentEmail
	check.Password = req.Password

	response, err := h.Auth.CheckEmailAndPassword(ctx, &check)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !response.Success {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect email or password"})
		return
	}

	code, err := h.SendEmail(new_email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send verification email"})
		return
	}
	req.Code = code

	res, err := h.Auth.ChangeEmail(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// VerifyUpdatedEmail godoc
// @Summary Verifies an updated email address
// @Description Confirms the updated email address with a verification code
// @Security BearerAuth
// @Tags Additional
// @Accept json
// @Produce json
// @Param new_email query string true "New Email"
// @Param code query string true "Verification code"
// @Success 200 {object} auth.InfoResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /additional/verify-new-email [post]
func (h *Handler) VerifyNewEmail(ctx *gin.Context) {
	new_email := ctx.Query("new_email")
	code := ctx.Query("code")

	id, err := token.GetIdFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	req := auth.VerifyNewEmailRequest{
		Id:       id,
		NewEmail: new_email,
		Code:     code,
	}

	res, err := h.Auth.VerifyNewEmail(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// RefreshToken godoc
// @Summary Refreshes a user's JWT token
// @Description Issues a new JWT token for the authenticated user
// @Security BearerAuth
// @Tags Additional
// @Accept json
// @Produce json
// @Success 200 {object} auth.TokenResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /additional/refresh-token [put]
func (h *Handler) RefreshToken(ctx *gin.Context) {
	id, err := token.GetIdFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	req := auth.ById{
		Id: id,
	}

	res, err := h.Auth.RefreshToken(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// ResetPasswordRequest godoc
// @Summary Requests a password reset for a user
// @Description Initiates a password reset process for the specified email address
// @Security BearerAuth
// @Tags Additional
// @Accept json
// @Produce json
// @Param email query string true "Email address to request a password reset for"
// @Success 200 {object} auth.InfoResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /additional/forget-password [post]
func (h *Handler) ForgetPassword(ctx *gin.Context) {
	email := ctx.Query("email")

	code, err := h.SendEmail(email)
	if err != nil {
		panic(err)
	}

	req := auth.ForgetPasswordRequest{
		Email: email,
		Code:  code,
	}

	res, err := h.Auth.ForgetPassword(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// ResetPassword godoc
// @Summary Resets a user's password
// @Description Resets the password for a user based on the provided details
// @Security BearerAuth
// @Tags Additional
// @Accept json
// @Produce json
// @Param email query string true "Email address to request a password reset for"
// @Param code query string true "Verification code"
// @Param new_password query string true "New Password"
// @Success 200 {object} auth.InfoResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /additional/reset-password [put]
func (h *Handler) ResetPassword(ctx *gin.Context) {
	email := ctx.Query("email")
	code := ctx.Query("code")
	new_password := ctx.Query("new_password")

	id, err := token.GetIdFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	req := auth.ResetPasswordRequest{
		Email:       email,
		Code:        code,
		NewPassword: new_password,
		Id:          id,
	}

	res, err := h.Auth.ResetPassword(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
