package product

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Router is the struct for handling product routes
type Router struct {
	handler *Handler
}

// NewRouter creates a new Router instance for the product module
func NewRouter(db *gorm.DB) *Router {
	// Initialize the service and handler for product
	service := &Service{repo: &Repository{DB: db}}
	handler := &Handler{service: service}
	return &Router{handler: handler}
}

// SetupRoutes registers product routes
func (r *Router) SetupRoutes(api *gin.RouterGroup) {
	productGroup := api.Group("/")
	{
		productGroup.GET("/", r.handler.GetAllProducts)
		productGroup.GET("/:id", r.handler.GetProductByID)
		productGroup.POST("/", r.handler.CreateProduct)
		productGroup.PUT("/:id", r.handler.UpdateProduct)
		productGroup.DELETE("/:id", r.handler.DeleteProduct)
	}
}
