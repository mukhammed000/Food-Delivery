package handler

import (
	"auth/api/token"
	"auth/genproto/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteUser godoc
// @Summary Deletes the user by entered admin's id
// @Description Deletes the user data
// @Security BearerAuth
// @Tags Admin
// @Accept json
// @Produce json
// @Param user_id query string true "User ID"
// @Success 200 {object} auth.InfoResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/delete-user [delete]
func (h *Handler) DeleteUser(ctx *gin.Context) {
	userID := ctx.Query("user_id")

	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		return
	}

	adminID, err := token.GetIdFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	req := auth.DeleteProfileByAdminReq{
		UserId:  userID,
		AdminId: adminID,
	}

	res, err := h.Auth.DeleteProfileByAdmin(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// ListAllUsers godoc
// @Summary Lists all users
// @Description Retrieves a list of all users
// @Security BearerAuth
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} auth.GetAllProfilesResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/list-users [get]
func (h *Handler) ListAllUsers(ctx *gin.Context) {
	adminID, err := token.GetIdFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	req := auth.ById{
		Id: adminID,
	}

	res, err := h.Auth.ListAllProfiles(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
