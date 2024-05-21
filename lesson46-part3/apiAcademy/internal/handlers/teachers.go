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

func (h *Handlers) GetAllTeachers(c *gin.Context) {
	var teachers []models.Teacher
	if err := h.DB.Preload("Groups.Course").
		Find(&teachers).Error; err != nil {
		log.Println("cannot get teachers:", err.Error())
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"teachers": teachers,
	})
}

type teacherRequest struct {
	FullName string `json:"fullName" binding:"required,max=64"`
	Subject  string `json:"subject" binding:"required,max=64"`
}

func (h *Handlers) CreateTeacher(c *gin.Context) {
	validationErrors := make(map[string]string)

	var requestBody teacherRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		helpers.FillValidationErrorTag(err, validationErrors)
	}

	if len(validationErrors) != 0 {
		c.JSON(http.StatusUnprocessableEntity, validationErrors)
		return
	}

	teacher := models.Teacher{
		FullName: requestBody.FullName,
		Subject:  requestBody.Subject,
	}

	if err := h.DB.Create(&teacher).Error; err != nil {
		log.Println("cannot create teacher:", err.Error())
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"teacher": teacher,
	})
}

func (h *Handlers) GetOneTeacher(c *gin.Context) {
	var teacher models.Teacher
	if err := h.DB.Where("id = ?", c.Param("id")).
		Preload("Groups.Course").
		First(&teacher).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.NotFound(c)
			return
		}
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"teacher": teacher,
	})
}

func (h *Handlers) UpdateTeacher(c *gin.Context) {
	var teacher models.Teacher
	if err := h.DB.Where("id = ?", c.Param("id")).
		First(&teacher).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.NotFound(c)
			return
		}
		helpers.InternalServerError(c)
		return
	}

	validationErrors := make(map[string]string)

	var requestBody teacherRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		helpers.FillValidationErrorTag(err, validationErrors)
	}

	if len(validationErrors) != 0 {
		c.JSON(http.StatusUnprocessableEntity, validationErrors)
		return
	}

	teacher.FullName = requestBody.FullName
	teacher.Subject = requestBody.Subject

	if err := h.DB.Save(&teacher).Error; err != nil {
		log.Println("cannot update teacher:", err.Error())
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"teacher": teacher,
	})
}

func (h *Handlers) DeleteTeacher(c *gin.Context) {
	var teacher models.Teacher
	if err := h.DB.Where("id = ?", c.Param("id")).
		First(&teacher).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.NotFound(c)
			return
		}
		helpers.InternalServerError(c)
		return
	}

	if err := h.DB.Delete(&teacher).Error; err != nil {
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Успешно удалено",
	})
}
