package order

import (
	"ddd/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service OrderService
}

func NewHandler(service OrderService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAllOrders(c *gin.Context) {
	orders, err := h.service.GetAllOrders()
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to fetch orders")
		return
	}
	utils.SendResponse(c, http.StatusOK, orders)
}

func (h *Handler) GetOrderByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid Order ID")
		return
	}
	order, err := h.service.GetOrderByID(uint(id))
	if err != nil {
		utils.SendError(c, http.StatusNotFound, "Order not found")
		return
	}
	utils.SendResponse(c, http.StatusOK, order)
}

func (h *Handler) CreateOrder(c *gin.Context) {
	var order Order
	if err := c.ShouldBindJSON(&order); err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.CreateOrder(&order); err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to create order")
		return
	}
	utils.SendResponse(c, http.StatusCreated, order)
}

func (h *Handler) UpdateOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid Order ID")
		return
	}
	var updatedOrder Order
	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.UpdateOrder(uint(id), &updatedOrder); err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to update order")
		return
	}
	utils.SendResponse(c, http.StatusOK, updatedOrder)
}

func (h *Handler) DeleteOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid Order ID")
		return
	}
	if err := h.service.DeleteOrder(uint(id)); err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to delete order")
		return
	}
	utils.SendResponse(c, http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
