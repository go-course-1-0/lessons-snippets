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

func (h *Handlers) GetAllCourses(c *gin.Context) {
	var courses []models.Course
	if err := h.DB.Preload("Groups.Teacher").
		Find(&courses).Error; err != nil {
		log.Println("cannot get courses:", err.Error())
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"courses": courses,
	})
}

type courseRequest struct {
	Title    string `json:"title" binding:"required,max=64"`
	Duration int    `json:"duration" binding:"required,gte=1,lte=36"`
}

func (h *Handlers) CreateCourse(c *gin.Context) {
	validationErrors := make(map[string]string)

	var requestBody courseRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		helpers.FillValidationErrorTag(err, validationErrors)
	}

	if len(validationErrors) != 0 {
		c.JSON(http.StatusUnprocessableEntity, validationErrors)
		return
	}

	course := models.Course{
		Title:    requestBody.Title,
		Duration: requestBody.Duration,
	}

	if err := h.DB.Create(&course).Error; err != nil {
		log.Println("cannot create course:", err.Error())
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"course": course,
	})
}

func (h *Handlers) GetOneCourse(c *gin.Context) {
	var course models.Course
	if err := h.DB.Where("id = ?", c.Param("id")).
		Preload("Groups.Teacher").
		First(&course).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.NotFound(c)
			return
		}
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"course": course,
	})
}

func (h *Handlers) UpdateCourse(c *gin.Context) {
	var course models.Course
	if err := h.DB.Where("id = ?", c.Param("id")).
		First(&course).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.NotFound(c)
			return
		}
		helpers.InternalServerError(c)
		return
	}

	validationErrors := make(map[string]string)

	var requestBody courseRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		helpers.FillValidationErrorTag(err, validationErrors)
	}

	if len(validationErrors) != 0 {
		c.JSON(http.StatusUnprocessableEntity, validationErrors)
		return
	}

	course.Title = requestBody.Title
	course.Duration = requestBody.Duration

	if err := h.DB.Save(&course).Error; err != nil {
		log.Println("cannot update course:", err.Error())
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"course": course,
	})
}

func (h *Handlers) DeleteCourse(c *gin.Context) {
	var course models.Course
	if err := h.DB.Where("id = ?", c.Param("id")).
		First(&course).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.NotFound(c)
			return
		}
		helpers.InternalServerError(c)
		return
	}

	if err := h.DB.Delete(&course).Error; err != nil {
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Успешно удалено",
	})
}
