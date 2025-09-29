package handlers

import (
	"net/http"
	"notification-hub/models"
	"notification-hub/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ContactGroupHandler struct {
	DB *gorm.DB
}

func NewContactGroupHandler(db *gorm.DB) *ContactGroupHandler {
	return &ContactGroupHandler{DB: db}
}

// Create contact group
func (h *ContactGroupHandler) Create(c *gin.Context) {
	var req models.ContactGroup
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Respond(c, http.StatusBadRequest, "Invalid request", nil, err)
		return
	}
	if err := h.DB.Create(&req).Error; err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Failed to create contact group", nil, err)
		return
	}
	utils.Respond(c, http.StatusCreated, "Contact group created successfully", req, nil)
}

// List contact groups
func (h *ContactGroupHandler) List(c *gin.Context) {
	var groups []models.ContactGroup
	if err := h.DB.Preload("Contacts").Find(&groups).Error; err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Failed to fetch contact groups", nil, err)
		return
	}
	utils.Respond(c, http.StatusOK, "Contact groups fetched successfully", groups, nil)
}
