package handler

import (
	"api/genproto/delivery"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateOffice handles the creation of a new office.
// @Summary Create a new office
// @Description Creates a new office based on the provided office details.
// @Tags Office
// @Accept json
// @Produce json
// @Param request body delivery.CreateOfficeRequest true "Create office request"
// @Success 200 {object} delivery.InfoResponse "Office created successfully"
// @Failure 400 {object} delivery.InfoResponse "Invalid request"
// @Failure 500 {object} delivery.InfoResponse "Internal server error"
// @Router /office [post]
func (h *Handler) CreateOffice(ctx *gin.Context) {
	var req delivery.CreateOfficeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, delivery.InfoResponse{Success: false, Message: "Invalid request"})
		return
	}

	resp, err := h.Office.CreateOffice(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, delivery.InfoResponse{Success: false, Message: "Failed to create office"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetOffice handles the retrieval of a specific office by its ID.
// @Summary Get office details
// @Description Retrieves details of an office based on the provided office ID.
// @Tags Office
// @Accept json
// @Produce json
// @Param office_id path string true "Office ID"
// @Success 200 {object} delivery.OfficeResponse "Office details retrieved successfully"
// @Failure 400 {object} delivery.InfoResponse "Invalid request"
// @Failure 500 {object} delivery.InfoResponse "Internal server error"
// @Router /office/{office_id} [get]
func (h *Handler) GetOffice(ctx *gin.Context) {
	officeID := ctx.Param("office_id")
	if officeID == "" {
		ctx.JSON(http.StatusBadRequest, delivery.InfoResponse{Success: false, Message: "Office ID is required"})
		return
	}

	resp, err := h.Office.GetOffice(ctx, &delivery.GetOfficeRequest{OfficeId: officeID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, delivery.InfoResponse{Success: false, Message: "Failed to retrieve office"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetAllOffice handles the retrieval of all offices with optional pagination.
// @Summary Get all offices
// @Description Retrieves a list of all offices with optional pagination.
// @Tags Office
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {object} delivery.GetAllOfficesResponse "List of all offices retrieved successfully"
// @Failure 400 {object} delivery.InfoResponse "Invalid request"
// @Failure 500 {object} delivery.InfoResponse "Internal server error"
// @Router /offices [get]
func (h *Handler) GetAllOffice(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page <= 0 {
		ctx.JSON(http.StatusBadRequest, delivery.InfoResponse{Success: false, Message: "Invalid page number"})
		return
	}

	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil || limit <= 0 {
		ctx.JSON(http.StatusBadRequest, delivery.InfoResponse{Success: false, Message: "Invalid limit number"})
		return
	}

	req := &delivery.GetAllOfficesRequest{
		Page:  int32(page),
		Limit: int32(page),
	}

	resp, err := h.Office.GetAllOffices(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, delivery.InfoResponse{Success: false, Message: "Failed to retrieve offices"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// UpdateOffice handles the update of an existing office.
// @Summary Update an existing office
// @Description Updates the details of an existing office based on the provided office ID and details.
// @Tags Office
// @Accept json
// @Produce json
// @Param request body delivery.UpdateOfficeRequest true "Update office request"
// @Success 200 {object} delivery.InfoResponse "Office updated successfully"
// @Failure 400 {object} delivery.InfoResponse "Invalid request"
// @Failure 404 {object} delivery.InfoResponse "Office not found"
// @Failure 500 {object} delivery.InfoResponse "Internal server error"
// @Router /office [put]
func (h *Handler) UpdateOffice(ctx *gin.Context) {
	var req delivery.UpdateOfficeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, delivery.InfoResponse{Success: false, Message: "Invalid request"})
		return
	}

	resp, err := h.Office.UpdateOffice(ctx, &req)
	if err != nil {
		if err.Error() == "office not found" {
			ctx.JSON(http.StatusNotFound, delivery.InfoResponse{Success: false, Message: "Office not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, delivery.InfoResponse{Success: false, Message: "Failed to update office"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// DeleteOffice handles the deletion of an office by its ID.
// @Summary Delete an office
// @Description Deletes an office based on the provided office ID.
// @Tags Office
// @Accept json
// @Produce json
// @Param request body delivery.DeleteOfficeRequest true "Delete office request"
// @Success 200 {object} delivery.InfoResponse "Office deleted successfully"
// @Failure 400 {object} delivery.InfoResponse "Invalid request"
// @Failure 404 {object} delivery.InfoResponse "Office not found"
// @Failure 500 {object} delivery.InfoResponse "Internal server error"
// @Router /office [delete]
func (h *Handler) DeleteOffice(ctx *gin.Context) {
	var req delivery.DeleteOfficeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, delivery.InfoResponse{Success: false, Message: "Invalid request"})
		return
	}

	resp, err := h.Office.DeleteOffice(ctx, &req)
	if err != nil {
		if err.Error() == "office not found" {
			ctx.JSON(http.StatusNotFound, delivery.InfoResponse{Success: false, Message: "Office not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, delivery.InfoResponse{Success: false, Message: "Failed to delete office"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
