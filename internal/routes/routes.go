package routes

import (
	"manage/internal/controllers"
	"manage/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes configures all API routes
func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Initialize controllers
	healthController := controllers.NewHealthController()
	churchController := controllers.NewChurchController(db)
	contactController := controllers.NewContactController(db)
	authController := controllers.NewAuthController(db)

	// API v1 routes
	api := r.Group("/api")
	{
		// Health check
		api.GET("/health", healthController.HealthCheck)
		
		// Public church routes
		api.GET("/home", churchController.GetHomeData)
		api.GET("/about", churchController.GetAboutData)
		api.GET("/services", churchController.GetAboutData) // Uses same data as about
		api.GET("/our-ministries", churchController.GetAboutData) // Uses same data as about
		api.GET("/events", churchController.GetEvents)
		api.GET("/gallery", churchController.GetContactData) // No specific data needed
		api.GET("/contact", churchController.GetContactData)
		api.GET("/give", churchController.GetContactData) // No specific data needed

		// Contact form routes
		api.POST("/contact/send", contactController.SendContactMessage)
		
		// Authentication routes (public)
		auth := api.Group("/auth")
		{
			auth.POST("/login", authController.Login)
			auth.POST("/register", authController.Register)
		}

		// Protected admin routes
		admin := api.Group("/admin")
		admin.Use(middleware.AuthMiddleware())
		{
			// Auth profile
			admin.GET("/profile", authController.GetProfile)
			admin.POST("/change-password", authController.ChangePassword)

			// Contact messages management
			admin.GET("/messages", contactController.GetContactMessages)
			admin.PATCH("/messages/:id/status", contactController.UpdateContactMessageStatus)
			admin.DELETE("/messages/:id", contactController.DeleteContactMessage)
		}
	}

	// Future API versions can be added here
	// v2 := r.Group("/api/v2")
}
