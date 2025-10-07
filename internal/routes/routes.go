package routes

import (
	"manage/internal/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes configures all API routes
func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Initialize controllers
	userController := controllers.NewUserController()
	healthController := controllers.NewHealthController()
	churchController := controllers.NewChurchController(db)

	// API v1 routes
	api := r.Group("/api")
	{
		// Health check
		api.GET("/health", healthController.HealthCheck)
		
		// User routes
		api.GET("/name", userController.GetName) // Used by frontend
		
		// Church routes
		api.GET("/home", churchController.GetHomeData)
		api.GET("/about", churchController.GetAboutData)
		api.GET("/services", churchController.GetAboutData) // Uses same data as about
		api.GET("/events", churchController.GetEvents)
		api.GET("/gallery", churchController.GetContactData) // No specific data needed
		api.GET("/contact", churchController.GetContactData)
		api.GET("/give", churchController.GetContactData) // No specific data needed
	}

	// Future API versions can be added here
	// v2 := r.Group("/api/v2")
}
