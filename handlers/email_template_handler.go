package handlers

import (
	"net/http"
	"notification-hub/models"
	"notification-hub/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EmailTemplateHandler struct {
	DB *gorm.DB
}

func NewEmailTemplateHandler(db *gorm.DB) *EmailTemplateHandler {
	return &EmailTemplateHandler{DB: db}
}

// Create Email Template
func (h *EmailTemplateHandler) Create(c *gin.Context) {
	var req models.EmailTemplate
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Respond(c, http.StatusBadRequest, "Invalid request", nil, err)
		return
	}
	if err := h.DB.Create(&req).Error; err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Failed to create email template", nil, err)
		return
	}
	utils.Respond(c, http.StatusCreated, "Email template created successfully", req, nil)
}

// List all Email Templates
func (h *EmailTemplateHandler) List(c *gin.Context) {
	var templates []models.EmailTemplate
	if err := h.DB.Find(&templates).Error; err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Failed to fetch email templates", nil, err)
		return
	}
	utils.Respond(c, http.StatusOK, "Email templates fetched successfully", templates, nil)
}

// Get Email Template by ID
func (h *EmailTemplateHandler) Get(c *gin.Context) {
	id := c.Param("id")
	var tmpl models.EmailTemplate
	if err := h.DB.First(&tmpl, "id = ?", id).Error; err != nil {
		utils.Respond(c, http.StatusNotFound, "Email template not found", nil, err)
		return
	}
	utils.Respond(c, http.StatusOK, "Email template fetched successfully", tmpl, nil)
}

// Update Email Template
func (h *EmailTemplateHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var tmpl models.EmailTemplate
	if err := h.DB.First(&tmpl, "id = ?", id).Error; err != nil {
		utils.Respond(c, http.StatusNotFound, "Email template not found", nil, err)
		return
	}

	var req models.EmailTemplate
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Respond(c, http.StatusBadRequest, "Invalid request", nil, err)
		return
	}

	tmpl.TemplateName = req.TemplateName
	tmpl.EmailTitle = req.EmailTitle
	tmpl.EmailBody = req.EmailBody
	tmpl.PropertyInfo = req.PropertyInfo
	tmpl.SenderName = req.SenderName
	tmpl.Signature = req.Signature
	tmpl.Active = req.Active
	tmpl.Html = req.Html

	if err := h.DB.Save(&tmpl).Error; err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Failed to update email template", nil, err)
		return
	}
	utils.Respond(c, http.StatusOK, "Email template updated successfully", tmpl, nil)
}

// Delete Email Template
func (h *EmailTemplateHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.DB.Delete(&models.EmailTemplate{}, "id = ?", id).Error; err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Failed to delete email template", nil, err)
		return
	}
	utils.Respond(c, http.StatusOK, "Email template deleted successfully", nil, nil)
}
