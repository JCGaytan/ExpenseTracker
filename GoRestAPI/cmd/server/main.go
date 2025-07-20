package main

import (
	"github.com/gin-gonic/gin"
	"expense-tracker/internal/config"
	"expense-tracker/internal/handlers"
	"expense-tracker/internal/middleware"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	_ "expense-tracker/docs"
)

// @title Expense Tracker API
// @version 1.0
// @description REST API for tracking personal expenses and incomes.
// @BasePath /
func main() {
	config.LoadConfig()
	config.RunMigrations(config.AppConfig.DBPath)
	config.InitDB()

	// Seed initial users for login testing
	config.SeedUsers()
	config.SeedDemoData()
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	// Swagger docs
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Auth routes
	r.POST("/auth/register", handlers.Register)
	r.POST("/auth/login", handlers.Login)

	// Protected routes
	authMiddleware := middleware.JWTAuthMiddleware()
	api := r.Group("", authMiddleware)

	// Transaction endpoints
	api.GET("/transactions", handlers.ListTransactions)
	api.POST("/transactions", handlers.CreateTransaction)
	api.GET("/transactions/:id", handlers.GetTransaction)
	api.PUT("/transactions/:id", handlers.UpdateTransaction)
	api.DELETE("/transactions/:id", handlers.DeleteTransaction)

	// Category endpoints
	api.GET("/categories", handlers.ListCategories)
	api.POST("/categories", handlers.CreateCategory)
	api.PUT("/categories/:id", handlers.UpdateCategory)
	api.DELETE("/categories/:id", handlers.DeleteCategory)

	api.GET("/reports/summary", handlers.GetSummary)

	r.Run()
}

