package product

import (
	"ddd/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service ProductService
}

func NewHandler(service ProductService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAllProducts(c *gin.Context) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to fetch products")
		return
	}
	utils.SendResponse(c, http.StatusOK, products)
}

func (h *Handler) GetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid Product ID")
		return
	}
	product, err := h.service.GetProductByID(uint(id))
	if err != nil {
		utils.SendError(c, http.StatusNotFound, "Product not found")
		return
	}
	utils.SendResponse(c, http.StatusOK, product)
}

func (h *Handler) CreateProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.CreateProduct(&product); err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to create product")
		return
	}
	utils.SendResponse(c, http.StatusCreated, product)
}

func (h *Handler) UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid Product ID")
		return
	}
	var updatedProduct Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.UpdateProduct(uint(id), &updatedProduct); err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to update product")
		return
	}
	utils.SendResponse(c, http.StatusOK, updatedProduct)
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid Product ID")
		return
	}
	if err := h.service.DeleteProduct(uint(id)); err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to delete product")
		return
	}
	utils.SendResponse(c, http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
