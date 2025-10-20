package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HealthController struct{
	db *gorm.DB
}

// NewHealthController creates a new health controller
func NewHealthController(db *gorm.DB) *HealthController {
	return &HealthController{db: db}
}

// HealthCheck returns the health status of the API
func (hc *HealthController) HealthCheck(c *gin.Context) {
	response := gin.H{
		"status": "ok",
		"service": "church-api",
		"database": "checking...",
	}

	// Check database connection
	if hc.db != nil {
		sqlDB, err := hc.db.DB()
		if err == nil {
			if err := sqlDB.Ping(); err == nil {
				response["database"] = "connected"
			} else {
				response["database"] = "disconnected"
			}
		} else {
			response["database"] = "error"
		}
	} else {
		response["database"] = "not_initialized"
	}

	c.JSON(http.StatusOK, response)
}
