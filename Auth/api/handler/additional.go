package handler

import (
	"auth/api/token"
	"auth/genproto/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ChangeUserPassword godoc
// @Summary Changes a user's password
// @Description Updates the user's password after validating the current password
// @Tags Additional
// @Accept json
// @Produce json
// @Param body body auth.ChangePasswordRequest true "User password change details"
// @Success 200 {object} auth.InfoResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/change-password [put]
func (h *Handler) ChangePassword(ctx *gin.Context) {
	var req auth.ChangePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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
// @Tags Additional
// @Accept json
// @Produce json
// @Param body body auth.VerifyEmailRequest true "Email verification details"
// @Success 200 {object} auth.TokenResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/confirmEmail [post]
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
// @Tags Additional
// @Accept json
// @Produce json
// @Param body body auth.ChangeEmailRequest true "Update email details"
// @Success 200 {object} auth.InfoResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/updateEmail [put]
func (h *Handler) ChangeEmail(ctx *gin.Context) {
	var req auth.ChangeEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
// @Tags Additional
// @Accept json
// @Produce json
// @Param body body auth.VerifyNewEmailRequest true "Updated email verification details"
// @Success 200 {object} auth.InfoResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/verifyUpdatedEmail [post]
func (h *Handler) VerifyNewEmail(ctx *gin.Context) {
	var req auth.VerifyNewEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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
// @Tags Additional
// @Accept json
// @Produce json
// @Success 200 {object} auth.TokenResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /auth/refreshToken [put]
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
// @Tags Additional
// @Accept json
// @Produce json
// @Param email query string true "Email address to request a password reset for"
// @Success 200 {object} auth.InfoResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /auth/request-password-reset [post]
func (h *Handler) ForgetPassword(ctx *gin.Context) {
	email := ctx.Query("email")

	req := auth.ForgetPasswordRequest{
		Email: email,
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
// @Tags Additional
// @Accept json
// @Produce json
// @Param body body auth.ResetPasswordRequest true "User password reset details"
// @Success 200 {object} auth.InfoResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /auth/resetPassword [put]
func (h *Handler) ResetPassword(ctx *gin.Context) {
	var req auth.ResetPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Auth.ResetPassword(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// GetUserProfile godoc
// @Summary Retrieves a user's profile
// @Description Fetches the user's profile information by their ID
// @Tags Additional
// @Accept json
// @Produce json
// @Param user_id query string true "User ID"
// @Success 200 {object} auth.GetProfileResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/getProfile [get]
func (h *Handler) GetProfile(ctx *gin.Context) {
	userID := ctx.Query("user_id")

	req := auth.ById{Id: userID}

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
// @Tags Additional
// @Accept json
// @Produce json
// @Param body body auth.UpdateProfileRequest true "Additional update details"
// @Success 200 {object} auth.InfoResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/updateProfile [put]
func (h *Handler) UpdateProfile(ctx *gin.Context) {
	var req auth.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Auth.UpdateProfile(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// RemoveProfile godoc
// @Summary Removes a user's profile
// @Description Deletes the user's profile based on their ID
// @Tags Additional
// @Accept json
// @Produce json
// @Param user_id query string true "User ID"
// @Success 200 {object} auth.InfoResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/removeProfile [delete]
func (h *Handler) DeleteProfile(ctx *gin.Context) {
	userID := ctx.Query("user_id")

	req := auth.ById{Id: userID}

	res, err := h.Auth.DeleteProfile(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
