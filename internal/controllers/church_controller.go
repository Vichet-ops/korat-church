package controllers

import (
	"net/http"
	"time"

	"manage/internal/models"

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
	// Get upcoming events
	var events []models.Event
	cc.db.Where("is_public = ? AND date >= ?", true, time.Now()).
		Order("date ASC").
		Limit(5).
		Find(&events)

	// Get church settings
	var settings models.ChurchSettings
	cc.db.First(&settings)

	// If no settings exist, create default
	if settings.ID == 0 {
		settings = models.ChurchSettings{
			ChurchName:    "Muang Thai Korat Church",
			ChurchAddress: "123 Main St, Nakhon Ratchasima, Thailand 30000",
			ChurchPhone:   "(555) 123-4567",
			ChurchEmail:   "info@muangthaikoratchurch.com",
			ServiceTimes:  `[{"day":"Sunday","time":"10:00 AM","service":"Main Worship Service"},{"day":"Wednesday","time":"7:00 PM","service":"Bible Study & Prayer"},{"day":"Friday","time":"6:30 PM","service":"Youth Activities & Fellowship"}]`,
		}
		cc.db.Create(&settings)
	}

	c.JSON(http.StatusOK, gin.H{
		"events":  events,
		"settings": settings,
	})
}

// GetAboutData returns data for the about page
func (cc *ChurchController) GetAboutData(c *gin.Context) {
	var settings models.ChurchSettings
	cc.db.First(&settings)

	// If no settings exist, create default
	if settings.ID == 0 {
		settings = models.ChurchSettings{
			ChurchName:    "Muang Thai Korat Church",
			ChurchAddress: "123 Main St, Nakhon Ratchasima, Thailand 30000",
			ChurchPhone:   "(555) 123-4567",
			ChurchEmail:   "info@muangthaikoratchurch.com",
			ServiceTimes:  `[{"day":"Sunday","time":"10:00 AM","service":"Main Worship Service"},{"day":"Wednesday","time":"7:00 PM","service":"Bible Study & Prayer"},{"day":"Friday","time":"6:30 PM","service":"Youth Activities & Fellowship"}]`,
		}
		cc.db.Create(&settings)
	}

	c.JSON(http.StatusOK, gin.H{
		"settings": settings,
	})
}

// GetEvents returns all public events
func (cc *ChurchController) GetEvents(c *gin.Context) {
	var events []models.Event
	cc.db.Where("is_public = ?", true).
		Order("date ASC").
		Find(&events)

	// Get church settings for footer
	var settings models.ChurchSettings
	cc.db.First(&settings)

	// If no settings exist, create default
	if settings.ID == 0 {
		settings = models.ChurchSettings{
			ChurchName:    "Muang Thai Korat Church",
			ChurchAddress: "123 Main St, Nakhon Ratchasima, Thailand 30000",
			ChurchPhone:   "(555) 123-4567",
			ChurchEmail:   "info@muangthaikoratchurch.com",
			ServiceTimes:  `[{"day":"Sunday","time":"10:00 AM","service":"Main Worship Service"},{"day":"Wednesday","time":"7:00 PM","service":"Bible Study & Prayer"},{"day":"Friday","time":"6:30 PM","service":"Youth Activities & Fellowship"}]`,
		}
		cc.db.Create(&settings)
	}

	c.JSON(http.StatusOK, gin.H{
		"events":   events,
		"settings": settings,
	})
}

// GetContactData returns contact information
func (cc *ChurchController) GetContactData(c *gin.Context) {
	var settings models.ChurchSettings
	cc.db.First(&settings)

	// If no settings exist, create default
	if settings.ID == 0 {
		settings = models.ChurchSettings{
			ChurchName:    "Muang Thai Korat Church",
			ChurchAddress: "123 Main St, Nakhon Ratchasima, Thailand 30000",
			ChurchPhone:   "(555) 123-4567",
			ChurchEmail:   "info@muangthaikoratchurch.com",
			ServiceTimes:  `[{"day":"Sunday","time":"10:00 AM","service":"Main Worship Service"},{"day":"Wednesday","time":"7:00 PM","service":"Bible Study & Prayer"},{"day":"Friday","time":"6:30 PM","service":"Youth Activities & Fellowship"}]`,
		}
		cc.db.Create(&settings)
	}

	c.JSON(http.StatusOK, gin.H{
		"settings": settings,
	})
}
