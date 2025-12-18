package handlers

import (
	"amass-test/models"
	"amass-test/responses"
	"amass-test/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service services.AuthService
}

func NewAuthHandler(service services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	data, err := h.service.Login(input.Username, input.Password)
	if err != nil {
		responses.ErrorResponse(c, http.StatusUnauthorized, err)
		return
	}

	responses.SuccessResponse(c, "Login successful", data)
}

func (h *AuthHandler) Register(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	user := models.User{
		Username: input.Username,
		Password: input.Password,
	}

	if err := h.service.Register(user); err != nil {
		responses.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	responses.SuccessResponse(c, "Registration successful", nil)
}
