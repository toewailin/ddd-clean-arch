// internal/modules/order/dto/order_request.go
package order

// CreateOrderRequest represents the data needed to create a new order
type CreateOrderRequest struct {
    UserID   string  `json:"user_id" validate:"required"` // User who placed the order (required)
    ProductID string `json:"product_id" validate:"required"` // Product being ordered (required)
    Quantity int     `json:"quantity" validate:"required,gt=0"` // Quantity of the product (required, must be > 0)
    Total    float64 `json:"total" validate:"required,gt=0"` // Total price for the order (required, must be > 0)
}

// UpdateOrderRequest represents the data needed to update an existing order
type UpdateOrderRequest struct {
    Quantity int     `json:"quantity,omitempty" validate:"omitempty,gt=0"` // Quantity to be updated (optional, must be > 0)
    Total    float64 `json:"total,omitempty" validate:"omitempty,gt=0"`    // Total price to be updated (optional, must be > 0)
}
