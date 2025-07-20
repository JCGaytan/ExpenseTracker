package handlers

import (
	"net/http"
	"strconv"
	"time"
	"expense-tracker/internal/models"
	"expense-tracker/internal/config"
	"github.com/gin-gonic/gin"
)

type TransactionInput struct {
	Amount      float64   `json:"amount" binding:"required"`
	Date        string    `json:"date" binding:"required"`
	CategoryID  uint      `json:"category_id" binding:"required"`
	Description string    `json:"description"`
}

// CreateTransaction creates a new transaction for the authenticated user
// @Summary Create transaction
// @Description Create a new transaction (income or expense) for the current user
// @Tags transactions
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param input body TransactionInput true "Transaction info"
// @Success 201 {object} models.Transaction
// @Failure 400 {object} gin.H{"error":string}
// @Failure 401 {object} gin.H{"error":string}
// @Router /transactions [post]
func CreateTransaction(c *gin.Context) {
	var input TransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := c.GetUint("user_id")
	parsedDate, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD."})
		return
	}
	tx := models.Transaction{
		Amount:      input.Amount,
		Date:        parsedDate,
		CategoryID:  input.CategoryID,
		UserID:      userID,
		Description: input.Description,
	}
	if err := config.DB.Create(&tx).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, tx)
}

// GetTransaction returns a transaction by ID for the authenticated user
// @Summary Get transaction
// @Description Get a transaction by ID for the current user
// @Tags transactions
// @Security BearerAuth
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} models.Transaction
// @Failure 401 {object} gin.H{"error":string}
// @Failure 404 {object} gin.H{"error":string}
// @Router /transactions/{id} [get]
func GetTransaction(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetUint("user_id")
	var tx models.Transaction
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&tx).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}
	c.JSON(http.StatusOK, tx)
}

// UpdateTransaction updates a transaction for the authenticated user
// @Summary Update transaction
// @Description Update a transaction for the current user
// @Tags transactions
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Param input body TransactionInput true "Transaction info"
// @Success 200 {object} models.Transaction
// @Failure 400 {object} gin.H{"error":string}
// @Failure 401 {object} gin.H{"error":string}
// @Failure 404 {object} gin.H{"error":string}
// @Router /transactions/{id} [put]
func UpdateTransaction(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetUint("user_id")
	var tx models.Transaction
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&tx).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}
	var input TransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	parsedDate, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD."})
		return
	}
	tx.Amount = input.Amount
	tx.Date = parsedDate
	tx.CategoryID = input.CategoryID
	tx.Description = input.Description
	config.DB.Save(&tx)
	c.JSON(http.StatusOK, tx)
}

// DeleteTransaction deletes a transaction for the authenticated user
// @Summary Delete transaction
// @Description Delete a transaction for the current user
// @Tags transactions
// @Security BearerAuth
// @Param id path int true "Transaction ID"
// @Success 204 {string} string ""
// @Failure 401 {object} gin.H{"error":string}
// @Failure 404 {object} gin.H{"error":string}
// @Router /transactions/{id} [delete]
func DeleteTransaction(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetUint("user_id")
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Transaction{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// ListTransactions returns all transactions for the authenticated user, with filters and pagination
// @Summary List transactions
// @Description Get all transactions for the current user, with optional filters and pagination
// @Tags transactions
// @Security BearerAuth
// @Produce json
// @Param start_date query string false "Start date (YYYY-MM-DD)"
// @Param end_date query string false "End date (YYYY-MM-DD)"
// @Param category_id query int false "Category ID"
// @Param min_amount query number false "Minimum amount"
// @Param max_amount query number false "Maximum amount"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} models.Transaction
// @Failure 401 {object} gin.H{"error":string}
// @Router /transactions [get]
func ListTransactions(c *gin.Context) {
	userID := c.GetUint("user_id")
	var txs []models.Transaction
	query := config.DB.Where("user_id = ?", userID)
	// Filtering (date, category, amount)
	if start := c.Query("start_date"); start != "" {
		query = query.Where("date >= ?", start)
	}
	if end := c.Query("end_date"); end != "" {
		query = query.Where("date <= ?", end)
	}
	if cat := c.Query("category_id"); cat != "" {
		query = query.Where("category_id = ?", cat)
	}
	if min := c.Query("min_amount"); min != "" {
		query = query.Where("amount >= ?", min)
	}
	if max := c.Query("max_amount"); max != "" {
		query = query.Where("amount <= ?", max)
	}
	// Pagination
	limit := 20
	offset := 0
	if l := c.Query("limit"); l != "" {
		if v, err := strconv.Atoi(l); err == nil {
			limit = v
		}
	}
	if o := c.Query("offset"); o != "" {
		if v, err := strconv.Atoi(o); err == nil {
			offset = v
		}
	}
	query = query.Limit(limit).Offset(offset)
	if err := query.Order("date desc").Find(&txs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, txs)
}
