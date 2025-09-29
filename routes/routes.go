package routes

import (
	"notification-hub/config"
	"notification-hub/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Context routes
	contextHandler := handlers.NewContextHandler(config.DB)
	contexts := r.Group("/contexts")
	{
		contexts.POST("", contextHandler.Create)
		contexts.GET("", contextHandler.List)
		contexts.GET("/:id", contextHandler.Get)
		contexts.PUT("/:id", contextHandler.Update)
		contexts.DELETE("/:id", contextHandler.Delete)
	}

	// Email Template routes
	emailTemplateHandler := handlers.NewEmailTemplateHandler(config.DB)
	templates := r.Group("/email-templates")
	{
		templates.POST("", emailTemplateHandler.Create)
		templates.GET("", emailTemplateHandler.List)
		templates.GET("/:id", emailTemplateHandler.Get)
		templates.PUT("/:id", emailTemplateHandler.Update)
		templates.DELETE("/:id", emailTemplateHandler.Delete)
	}

	// Group routes
	groupHandler := handlers.NewGroupHandler(config.DB)
	groups := r.Group("/groups")
	{
		groups.POST("", groupHandler.Create)
		groups.GET("", groupHandler.List)
		groups.GET("/:id", groupHandler.Get)
		groups.PUT("/:id", groupHandler.Update)
		groups.DELETE("/:id", groupHandler.Delete)
	}

	// Contact
	contactHandler := handlers.NewContactHandler(config.DB)
	r.POST("/contacts", contactHandler.Create)
	r.GET("/contacts", contactHandler.List)

	// Contact Group
	contactGroupHandler := handlers.NewContactGroupHandler(config.DB)
	r.POST("/contact-groups", contactGroupHandler.Create)
	r.GET("/contact-groups", contactGroupHandler.List)
}
