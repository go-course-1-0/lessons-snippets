package handlers

import (
	"apiAcademy/internal/database/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handlers) GetAllAdmins(c *gin.Context) {
	var admins []models.Admin
	if err := h.DB.Find(&admins).Error; err != nil {
		log.Println("cannot get admins:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"admins": admins,
	})
}

type adminRequest struct {
	FullName string `json:"fullName" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (h *Handlers) CreateAdmin(c *gin.Context) {
	var requestBody adminRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	admin := models.Admin{
		FullName: requestBody.FullName,
		Email:    requestBody.Email,
		Password: requestBody.Password,
	}

	if err := h.DB.Create(&admin).Error; err != nil {
		log.Println("cannot create admin:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, admin)
}

func (h *Handlers) GetOneAdmin(c *gin.Context) {
}

func (h *Handlers) UpdateAdmin(c *gin.Context) {
}

func (h *Handlers) DeleteAdmin(c *gin.Context) {
}
