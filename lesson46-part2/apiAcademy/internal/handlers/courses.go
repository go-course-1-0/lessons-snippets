package handlers

import (
	"apiAcademy/internal/database/models"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetAllCourses(c *gin.Context) {
	var admins []models.Admin
	h.DB.Find(&admins)

	c.JSON(200, admins)
}

func (h *Handlers) CreateCourse(c *gin.Context) {
	var admin models.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(422, "Validation error")
		return
	}

	h.DB.Create(&admin)

	c.JSON(201, admin)
}

func (h *Handlers) GetOneCourse(c *gin.Context) {
}

func (h *Handlers) UpdateCourse(c *gin.Context) {
}

func (h *Handlers) DeleteCourse(c *gin.Context) {
}
