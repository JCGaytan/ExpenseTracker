package handlers

import (
	"net/http"
	"expense-tracker/internal/config"
	"github.com/gin-gonic/gin"
)

type SummaryResponse struct {
	TotalIncome   float64            `json:"total_income"`
	TotalExpense  float64            `json:"total_expense"`
	ByCategory    map[string]float64 `json:"by_category"`
}

// GetSummary returns monthly totals and category breakdown for the authenticated user
// @Summary Get monthly totals and category breakdown
// @Description Returns total income, total expense, and a breakdown by category for the current user
// @Tags reports
// @Security BearerAuth
// @Produce json
// @Success 200 {object} SummaryResponse
// @Failure 401 {object} gin.H{"error":string}
// @Router /reports/summary [get]
func GetSummary(c *gin.Context) {
	userID := c.GetUint("user_id")
	var income, expense float64
	config.DB.Table("transactions").Where("user_id = ? AND amount > 0", userID).Select("SUM(amount)").Scan(&income)
	config.DB.Table("transactions").Where("user_id = ? AND amount < 0", userID).Select("SUM(amount)").Scan(&expense)

	// Category breakdown
	rows, err := config.DB.Raw(`SELECT c.name, SUM(t.amount) FROM transactions t JOIN categories c ON t.category_id = c.id WHERE t.user_id = ? GROUP BY c.name`, userID).Rows()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	byCat := make(map[string]float64)
	for rows.Next() {
		var name string
		var sum float64
		rows.Scan(&name, &sum)
		byCat[name] = sum
	}
	c.JSON(http.StatusOK, SummaryResponse{
		TotalIncome:  income,
		TotalExpense: expense,
		ByCategory:   byCat,
	})
}
