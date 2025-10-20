package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ChurchController struct {
	db *gorm.DB
}

func NewChurchController(db *gorm.DB) *ChurchController {
	return &ChurchController{db: db}
}

// GetHomeData returns data for the home page
func (cc *ChurchController) GetHomeData(c *gin.Context) {
	// Return static church settings for now
	settings := map[string]interface{}{
		"church_name":    "Muang Thai Korat Church",
		"church_address": "123 Main St, Nakhon Ratchasima, Thailand 30000",
		"church_phone":   "(555) 123-4567",
		"church_email":   "info@muangthaikoratchurch.com",
		"service_times":  `[{"day":"Sunday","time":"10:00 AM","service":"Main Worship Service"},{"day":"Wednesday","time":"7:00 PM","service":"Bible Study & Prayer"},{"day":"Friday","time":"6:30 PM","service":"Youth Activities & Fellowship"}]`,
	}

	c.JSON(http.StatusOK, gin.H{
		"events":   []interface{}{},
		"settings": settings,
	})
}

// GetAboutData returns data for the about page
func (cc *ChurchController) GetAboutData(c *gin.Context) {
	settings := map[string]interface{}{
		"church_name":    "Muang Thai Korat Church",
		"church_address": "123 Main St, Nakhon Ratchasima, Thailand 30000",
		"church_phone":   "(555) 123-4567",
		"church_email":   "info@muangthaikoratchurch.com",
		"service_times":  `[{"day":"Sunday","time":"10:00 AM","service":"Main Worship Service"},{"day":"Wednesday","time":"7:00 PM","service":"Bible Study & Prayer"},{"day":"Friday","time":"6:30 PM","service":"Youth Activities & Fellowship"}]`,
	}

	c.JSON(http.StatusOK, gin.H{
		"settings": settings,
	})
}

// GetEvents returns all public events
func (cc *ChurchController) GetEvents(c *gin.Context) {
	settings := map[string]interface{}{
		"church_name":    "Muang Thai Korat Church",
		"church_address": "123 Main St, Nakhon Ratchasima, Thailand 30000",
		"church_phone":   "(555) 123-4567",
		"church_email":   "info@muangthaikoratchurch.com",
		"service_times":  `[{"day":"Sunday","time":"10:00 AM","service":"Main Worship Service"},{"day":"Wednesday","time":"7:00 PM","service":"Bible Study & Prayer"},{"day":"Friday","time":"6:30 PM","service":"Youth Activities & Fellowship"}]`,
	}

	c.JSON(http.StatusOK, gin.H{
		"events":   []interface{}{},
		"settings": settings,
	})
}

// GetContactData returns contact information
func (cc *ChurchController) GetContactData(c *gin.Context) {
	settings := map[string]interface{}{
		"church_name":    "Muang Thai Korat Church",
		"church_address": "123 Main St, Nakhon Ratchasima, Thailand 30000",
		"church_phone":   "(555) 123-4567",
		"church_email":   "info@muangthaikoratchurch.com",
		"service_times":  `[{"day":"Sunday","time":"10:00 AM","service":"Main Worship Service"},{"day":"Wednesday","time":"7:00 PM","service":"Bible Study & Prayer"},{"day":"Friday","time":"6:30 PM","service":"Youth Activities & Fellowship"}]`,
	}

	c.JSON(http.StatusOK, gin.H{
		"settings": settings,
	})
}
