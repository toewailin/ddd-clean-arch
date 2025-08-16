package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Router is the struct for handling user routes
type Router struct {
	handler *Handler
}

// NewRouter creates a new Router instance for the user module
func NewRouter(db *gorm.DB) *Router {
	// Initialize the service and handler for user
	service := &Service{repo: &Repository{DB: db}}
	handler := &Handler{service: service}
	return &Router{handler: handler}
}

// SetupRoutes registers user routes
func (r *Router) SetupRoutes(api *gin.RouterGroup) {
	// Remove the redundant "users" path here
	// We already have "/users" in the main router, so no need to add "/users" again.
	userGroup := api.Group("") // <-- Just leave it as an empty string to avoid duplication
	{
		userGroup.GET("/", r.handler.GetAllUsers)       // List all users
		userGroup.GET("/:id", r.handler.GetUserByID)   // Get a user by ID
		userGroup.POST("/", r.handler.CreateUser)      // Create a new user
		userGroup.PUT("/:id", r.handler.UpdateUser)    // Update a user by ID
		userGroup.DELETE("/:id", r.handler.DeleteUser) // Delete a user by ID
	}
}
