package handlers

import (
	"net/http"
	"strconv"
	"expense-tracker/internal/models"
	"expense-tracker/internal/config"
	"github.com/gin-gonic/gin"
)

type CategoryInput struct {
	Name string `json:"name" binding:"required"`
}

// ListCategories returns all categories for the authenticated user
// @Summary List categories
// @Description Get all categories for the current user
// @Tags categories
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.Category
// @Failure 401 {object} gin.H{"error":string}
// @Router /categories [get]
func ListCategories(c *gin.Context) {
	userID := c.GetUint("user_id")
	var cats []models.Category
	if err := config.DB.Where("user_id = ?", userID).Find(&cats).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cats)
}

// CreateCategory creates a new category for the authenticated user
// @Summary Create category
// @Description Create a new category for the current user
// @Tags categories
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param input body CategoryInput true "Category info"
// @Success 201 {object} models.Category
// @Failure 400 {object} gin.H{"error":string}
// @Failure 401 {object} gin.H{"error":string}
// @Router /categories [post]
func CreateCategory(c *gin.Context) {
	userID := c.GetUint("user_id")
	var input CategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cat := models.Category{Name: input.Name, UserID: userID}
	if err := config.DB.Create(&cat).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, cat)
}

// UpdateCategory updates a category for the authenticated user
// @Summary Update category
// @Description Update a category for the current user
// @Tags categories
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param input body CategoryInput true "Category info"
// @Success 200 {object} models.Category
// @Failure 400 {object} gin.H{"error":string}
// @Failure 401 {object} gin.H{"error":string}
// @Failure 404 {object} gin.H{"error":string}
// @Router /categories/{id} [put]
func UpdateCategory(c *gin.Context) {
	userID := c.GetUint("user_id")
	id, _ := strconv.Atoi(c.Param("id"))
	var cat models.Category
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&cat).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	var input CategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cat.Name = input.Name
	config.DB.Save(&cat)
	c.JSON(http.StatusOK, cat)
}

// DeleteCategory deletes a category for the authenticated user
// @Summary Delete category
// @Description Delete a category for the current user
// @Tags categories
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Success 204 {string} string ""
// @Failure 401 {object} gin.H{"error":string}
// @Failure 404 {object} gin.H{"error":string}
// @Router /categories/{id} [delete]
func DeleteCategory(c *gin.Context) {
	userID := c.GetUint("user_id")
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Category{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
