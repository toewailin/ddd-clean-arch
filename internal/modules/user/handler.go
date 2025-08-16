package user

import (
	"ddd/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service UserService
}

func NewHandler(service UserService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to fetch users")
		return
	}
	utils.SendResponse(c, http.StatusOK, users)
}

func (h *Handler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid User ID")
		return
	}
	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		utils.SendError(c, http.StatusNotFound, "User not found")
		return
	}
	utils.SendResponse(c, http.StatusOK, user)
}

func (h *Handler) CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.CreateUser(&user); err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to create user")
		return
	}
	utils.SendResponse(c, http.StatusCreated, user)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid User ID")
		return
	}
	var updatedUser User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.UpdateUser(uint(id), &updatedUser); err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to update user")
		return
	}
	utils.SendResponse(c, http.StatusOK, updatedUser)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid User ID")
		return
	}
	if err := h.service.DeleteUser(uint(id)); err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to delete user")
		return
	}
	utils.SendResponse(c, http.StatusOK, gin.H{"message": "User deleted successfully"})
}
