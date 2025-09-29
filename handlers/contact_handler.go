package handlers

import (
	"net/http"
	"notification-hub/models"
	"notification-hub/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ContactHandler struct {
	DB *gorm.DB
}

func NewContactHandler(db *gorm.DB) *ContactHandler {
	return &ContactHandler{DB: db}
}

// Create contact
func (h *ContactHandler) Create(c *gin.Context) {
	var req models.Contact
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Respond(c, http.StatusBadRequest, "Invalid request", nil, err)
		return
	}
	if err := h.DB.Create(&req).Error; err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Failed to create contact", nil, err)
		return
	}
	utils.Respond(c, http.StatusCreated, "Contact created successfully", req, nil)
}

// List contacts
func (h *ContactHandler) List(c *gin.Context) {
	var contacts []models.Contact
	if err := h.DB.Find(&contacts).Error; err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Failed to fetch contacts", nil, err)
		return
	}
	utils.Respond(c, http.StatusOK, "Contacts fetched successfully", contacts, nil)
}
