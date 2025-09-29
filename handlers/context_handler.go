package handlers

import (
	"net/http"
	"notification-hub/models"
	"notification-hub/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ContextHandler struct {
	DB *gorm.DB
}

func NewContextHandler(db *gorm.DB) *ContextHandler {
	return &ContextHandler{DB: db}
}

func (h *ContextHandler) Create(c *gin.Context) {
	var req models.NotificationContext
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Respond(c, http.StatusBadRequest, "Invalid request", nil, err)
		return
	}
	if err := h.DB.Create(&req).Error; err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Failed to create context", nil, err)
		return
	}
	utils.Respond(c, http.StatusCreated, "Context created successfully", req, nil)
}

func (h *ContextHandler) List(c *gin.Context) {
	var contexts []models.NotificationContext
	if err := h.DB.Find(&contexts).Error; err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Failed to fetch contexts", nil, err)
		return
	}
	utils.Respond(c, http.StatusOK, "Contexts fetched successfully", contexts, nil)
}

func (h *ContextHandler) Get(c *gin.Context) {
	id := c.Param("id")
	var ctx models.NotificationContext
	if err := h.DB.First(&ctx, "id = ?", id).Error; err != nil {
		utils.Respond(c, http.StatusNotFound, "Context not found", nil, err)
		return
	}
	utils.Respond(c, http.StatusOK, "Context fetched successfully", ctx, nil)
}

func (h *ContextHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var ctx models.NotificationContext
	if err := h.DB.First(&ctx, "id = ?", id).Error; err != nil {
		utils.Respond(c, http.StatusNotFound, "Context not found", nil, err)
		return
	}

	var req models.NotificationContext
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Respond(c, http.StatusBadRequest, "Invalid request", nil, err)
		return
	}

	ctx.Name = req.Name
	ctx.Description = req.Description
	if err := h.DB.Save(&ctx).Error; err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Failed to update context", nil, err)
		return
	}
	utils.Respond(c, http.StatusOK, "Context updated successfully", ctx, nil)
}

func (h *ContextHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.DB.Delete(&models.NotificationContext{}, "id = ?", id).Error; err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Failed to delete context", nil, err)
		return
	}
	utils.Respond(c, http.StatusOK, "Context deleted successfully", nil, nil)
}
