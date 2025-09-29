package handlers

import (
	"net/http"
	"notification-hub/models"
	"notification-hub/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GroupHandler struct {
	DB *gorm.DB
}

func NewGroupHandler(db *gorm.DB) *GroupHandler {
	return &GroupHandler{DB: db}
}

// Create
func (h *GroupHandler) Create(c *gin.Context) {
	var req models.Group
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Respond(c, http.StatusBadRequest, "Invalid request", nil, err)
		return
	}
	if err := h.DB.Create(&req).Error; err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Failed to create group", nil, err)
		return
	}
	utils.Respond(c, http.StatusCreated, "Group created successfully", req, nil)
}

// List
func (h *GroupHandler) List(c *gin.Context) {
	var groups []models.Group
	if err := h.DB.Find(&groups).Error; err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Failed to fetch groups", nil, err)
		return
	}
	utils.Respond(c, http.StatusOK, "Groups fetched successfully", groups, nil)
}

// Get
func (h *GroupHandler) Get(c *gin.Context) {
	id := c.Param("id")
	var group models.Group
	if err := h.DB.First(&group, "id = ?", id).Error; err != nil {
		utils.Respond(c, http.StatusNotFound, "Group not found", nil, err)
		return
	}
	utils.Respond(c, http.StatusOK, "Group fetched successfully", group, nil)
}

// Update
func (h *GroupHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var group models.Group
	if err := h.DB.First(&group, "id = ?", id).Error; err != nil {
		utils.Respond(c, http.StatusNotFound, "Group not found", nil, err)
		return
	}

	var req models.Group
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Respond(c, http.StatusBadRequest, "Invalid request", nil, err)
		return
	}

	group.GroupName = req.GroupName
	group.NotificationEvent = req.NotificationEvent
	group.Site = req.Site
	group.Description = req.Description
	group.Channels = req.Channels
	group.Contacts = req.Contacts
	group.ContactGroups = req.ContactGroups
	group.Active = req.Active

	if err := h.DB.Save(&group).Error; err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Failed to update group", nil, err)
		return
	}
	utils.Respond(c, http.StatusOK, "Group updated successfully", group, nil)
}

// Delete
func (h *GroupHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.DB.Delete(&models.Group{}, "id = ?", id).Error; err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Failed to delete group", nil, err)
		return
	}
	utils.Respond(c, http.StatusOK, "Group deleted successfully", nil, nil)
}
