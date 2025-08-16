package faq

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Router is the struct for handling FAQ routes
type Router struct {
	handler *Handler
}

// NewRouter creates a new Router instance for the FAQ module
func NewRouter(db *gorm.DB) *Router {
	// Initialize the service and handler for FAQ
	repo := &Repository{DB: db}
	service := &Service{repo: repo}
	handler := &Handler{service: service}
	return &Router{handler: handler}
}

// SetupRoutes registers FAQ routes
func (r *Router) SetupRoutes(rg *gin.RouterGroup) {
	// Register routes for FAQ
	faqGroup := rg.Group("")
	{
		faqGroup.GET("/", r.handler.GetAllFAQs)     // List all FAQs
		faqGroup.GET("/:id", r.handler.GetFAQByID)  // Get FAQ by ID
		faqGroup.POST("/", r.handler.CreateFAQ)     // Create a new FAQ
		faqGroup.PUT("/:id", r.handler.UpdateFAQ)   // Update FAQ by ID
		faqGroup.DELETE("/:id", r.handler.DeleteFAQ) // Delete FAQ by ID
	}
}
