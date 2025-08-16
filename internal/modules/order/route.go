package order

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Router is the struct for handling order routes
type Router struct {
	handler *Handler
}

// NewRouter creates a new Router instance for the order module
func NewRouter(db *gorm.DB) *Router {
	// Initialize the service and handler for order
	service := &Service{repo: &Repository{DB: db}}
	handler := &Handler{service: service}
	return &Router{handler: handler}
}

// SetupRoutes registers order routes
func (r *Router) SetupRoutes(api *gin.RouterGroup) {
	orderGroup := api.Group("")
	{
		orderGroup.GET("/", r.handler.GetAllOrders)
		orderGroup.GET("/:id", r.handler.GetOrderByID)
		orderGroup.POST("/", r.handler.CreateOrder)
		orderGroup.PUT("/:id", r.handler.UpdateOrder)
		orderGroup.DELETE("/:id", r.handler.DeleteOrder)
	}
}
