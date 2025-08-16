package router

import (
	"ddd/internal/modules/faq"
	"ddd/internal/modules/order"
	"ddd/internal/modules/product"
	"ddd/internal/modules/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes initializes the routes and injects dependencies
func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	// Initialize routers for each module with their respective dependencies
	faqRouter := faq.NewRouter(db)
	productRouter := product.NewRouter(db)
	orderRouter := order.NewRouter(db)
	userRouter := user.NewRouter(db) // New router for user module

	// API group
	api := r.Group("/api")
	{
		// FAQ Routes
		faqGroup := api.Group("/faqs")
		faqRouter.SetupRoutes(faqGroup)

		// Product Routes
		productGroup := api.Group("/products")
		productRouter.SetupRoutes(productGroup)

		// Order Routes
		orderGroup := api.Group("/orders")
		orderRouter.SetupRoutes(orderGroup)

		// User Routes
		userGroup := api.Group("/users")
		userRouter.SetupRoutes(userGroup) // Register user routes
	}

	// You can add more modules here as your application grows
}
