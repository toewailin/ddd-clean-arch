package faq

import (
	"ddd/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service FAQService
}

func NewHandler(service FAQService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAllFAQs(c *gin.Context) {
	faqs, err := h.service.GetAllFAQs()
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to fetch FAQs")
		return
	}
	utils.SendResponse(c, http.StatusOK, faqs)
}

func (h *Handler) GetFAQByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid FAQ ID")
		return
	}
	faq, err := h.service.GetFAQByID(uint(id))
	if err != nil {
		utils.SendError(c, http.StatusNotFound, "FAQ not found")
		return
	}
	utils.SendResponse(c, http.StatusOK, faq)
}

func (h *Handler) CreateFAQ(c *gin.Context) {
	var faq FAQ
	if err := c.ShouldBindJSON(&faq); err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.CreateFAQ(&faq); err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to create FAQ")
		return
	}
	utils.SendResponse(c, http.StatusCreated, faq)
}

func (h *Handler) UpdateFAQ(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid FAQ ID")
		return
	}
	var updatedFAQ FAQ
	if err := c.ShouldBindJSON(&updatedFAQ); err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.UpdateFAQ(uint(id), &updatedFAQ); err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to update FAQ")
		return
	}
	utils.SendResponse(c, http.StatusOK, updatedFAQ)
}

func (h *Handler) DeleteFAQ(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid FAQ ID")
		return
	}
	if err := h.service.DeleteFAQ(uint(id)); err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to delete FAQ")
		return
	}
	utils.SendResponse(c, http.StatusOK, gin.H{"message": "FAQ deleted successfully"})
}