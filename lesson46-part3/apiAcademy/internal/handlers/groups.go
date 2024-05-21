package handlers

import (
	"apiAcademy/internal/database/models"
	"apiAcademy/internal/helpers"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func (h *Handlers) GetAllGroups(c *gin.Context) {
	var groups []models.Group
	if err := h.DB.Preload("Course").
		Preload("Teacher").
		Find(&groups).Error; err != nil {
		log.Println("cannot get groups:", err.Error())
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"groups": groups,
	})
}

// UUID or GUID
type groupRequest struct {
	CourseID  int       `json:"courseID" binding:"required,gte=1"`
	TeacherID int       `json:"teacherID" binding:"required,gte=1"`
	Title     string    `json:"title" binding:"required,max=64"`
	Start     time.Time `json:"start" binding:"omitempty,datetime=2006-01-02"`
	Finish    time.Time `json:"finish" binding:"omitempty,datetime=2006-01-02"`
}

func (h *Handlers) CreateGroup(c *gin.Context) {
	validationErrors := make(map[string]string)

	var requestBody groupRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		helpers.FillValidationErrorTag(err, validationErrors)
	}

	if len(validationErrors) != 0 {
		c.JSON(http.StatusUnprocessableEntity, validationErrors)
		return
	}

	group := models.Group{
		CourseID:  requestBody.CourseID,
		TeacherID: requestBody.TeacherID,
		Title:     requestBody.Title,
		Start:     requestBody.Start,
		Finish:    requestBody.Finish,
	}

	if err := h.DB.Create(&group).Error; err != nil {
		log.Println("cannot create group:", err.Error())
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"group": group,
	})
}

func (h *Handlers) GetOneGroup(c *gin.Context) {
	var group models.Group
	if err := h.DB.Where("id = ?", c.Param("id")).
		Preload("Course").
		Preload("Teacher").
		Preload("Students").
		Preload("Lessons").
		First(&group).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.NotFound(c)
			return
		}
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"group": group,
	})
}

func (h *Handlers) UpdateGroup(c *gin.Context) {
	var group models.Group
	if err := h.DB.Where("id = ?", c.Param("id")).
		First(&group).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.NotFound(c)
			return
		}
		helpers.InternalServerError(c)
		return
	}

	validationErrors := make(map[string]string)

	var requestBody groupRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		helpers.FillValidationErrorTag(err, validationErrors)
	}

	if len(validationErrors) != 0 {
		c.JSON(http.StatusUnprocessableEntity, validationErrors)
		return
	}

	group.CourseID = requestBody.CourseID
	group.TeacherID = requestBody.TeacherID
	group.Title = requestBody.Title
	group.Start = requestBody.Start
	group.Finish = requestBody.Finish

	if err := h.DB.Save(&group).Error; err != nil {
		log.Println("cannot update group:", err.Error())
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"group": group,
	})
}

func (h *Handlers) DeleteGroup(c *gin.Context) {
	var group models.Group
	if err := h.DB.Where("id = ?", c.Param("id")).
		First(&group).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.NotFound(c)
			return
		}
		helpers.InternalServerError(c)
		return
	}

	if err := h.DB.Delete(&group).Error; err != nil {
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Успешно удалено",
	})
}
