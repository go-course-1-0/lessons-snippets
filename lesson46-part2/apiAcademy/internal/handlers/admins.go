package handlers

import (
	"apiAcademy/internal/database/models"
	"apiAcademy/internal/helpers"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func (h *Handlers) GetAllAdmins(c *gin.Context) {
	var admins []models.Admin
	if err := h.DB.Find(&admins).Error; err != nil {
		log.Println("cannot get admins:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"admins": admins,
	})
}

type adminRequest struct {
	FullName string `json:"fullName" binding:"required,max=64"`
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

func (h *Handlers) CreateAdmin(c *gin.Context) {
	validationErrors := make(map[string]string)

	var requestBody adminRequest
	// data validation
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		helpers.FillValidationErrorTag(err, validationErrors)
	}

	// business validation
	//select count(*) from admins where email='email' // if count == 0
	//select * from admins where email='email' limit 1 // if null
	if err := h.DB.Where("email = ?", requestBody.Email).
		First(&models.Admin{}).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		if _, exists := validationErrors["email"]; !exists {
			validationErrors["email"] = helpers.ValidationMessageForTag("unique", "")
		}
	}

	if len(validationErrors) != 0 {
		c.JSON(http.StatusUnprocessableEntity, validationErrors)
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
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"admin": admin,
	})
}

func (h *Handlers) GetOneAdmin(c *gin.Context) {
	var admin models.Admin
	if err := h.DB.Where("id = ?", c.Param("id")).
		First(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Record Not Found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"admin": admin,
	})
}

func (h *Handlers) UpdateAdmin(c *gin.Context) {
	var admin models.Admin
	if err := h.DB.Where("id = ?", c.Param("id")).
		First(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Record Not Found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	validationErrors := make(map[string]string)

	var requestBody adminRequest
	// data validation
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		helpers.FillValidationErrorTag(err, validationErrors)
	}

	// business validation
	//select count(*) from admins where email='email' // if count == 0
	//select * from admins where email='email' limit 1 // if null
	if err := h.DB.Where("email = ?", requestBody.Email).
		Where("id != ?", admin.ID).
		First(&models.Admin{}).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		if _, exists := validationErrors["email"]; !exists {
			validationErrors["email"] = helpers.ValidationMessageForTag("unique", "")
		}
	}

	if len(validationErrors) != 0 {
		c.JSON(http.StatusUnprocessableEntity, validationErrors)
		return
	}

	admin.FullName = requestBody.FullName
	admin.Email = requestBody.Email
	admin.Password = requestBody.Password

	if err := h.DB.Save(&admin).Error; err != nil {
		log.Println("cannot update admin:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"admin": admin,
	})
}

func (h *Handlers) DeleteAdmin(c *gin.Context) {
	var admin models.Admin
	if err := h.DB.Where("id = ?", c.Param("id")).
		First(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Record Not Found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	// delete from admins where id = admin.ID
	if err := h.DB.Delete(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Администратор удален",
	})
}
