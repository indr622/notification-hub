package handlers

import (
	"net/http"
	"notification-hub/models"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, req)
}

func (h *ContextHandler) List(c *gin.Context) {
	var contexts []models.NotificationContext
	if err := h.DB.Find(&contexts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contexts)
}

func (h *ContextHandler) Get(c *gin.Context) {
	id := c.Param("id")
	var ctx models.NotificationContext
	if err := h.DB.First(&ctx, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Context not found"})
		return
	}
	c.JSON(http.StatusOK, ctx)
}

func (h *ContextHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var ctx models.NotificationContext
	if err := h.DB.First(&ctx, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Context not found"})
		return
	}

	var req models.NotificationContext
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Name = req.Name
	ctx.Description = req.Description
	if err := h.DB.Save(&ctx).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ctx)
}

func (h *ContextHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.DB.Delete(&models.NotificationContext{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
