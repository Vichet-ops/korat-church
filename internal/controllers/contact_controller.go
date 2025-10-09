package controllers

import (
	"net/http"

	"manage/internal/models"
	"manage/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ContactController struct {
	db           *gorm.DB
	emailService *services.EmailService
}

func NewContactController(db *gorm.DB) *ContactController {
	return &ContactController{
		db:           db,
		emailService: services.NewEmailService(),
	}
}

// SendContactMessage handles contact form submissions
func (cc *ContactController) SendContactMessage(c *gin.Context) {
	var message models.ContactMessage

	// Bind and validate JSON input
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input",
			"details": err.Error(),
		})
		return
	}

	// Save message to database
	if err := cc.db.Create(&message).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save message",
		})
		return
	}

	// Send email notification to admin (non-blocking)
	go func() {
		emailData := services.ContactEmailData{
			Name:    message.Name,
			Email:   message.Email,
			Subject: message.Subject,
			Message: message.Message,
		}
		
		// Send notification to admin
		if err := cc.emailService.SendContactNotification(emailData); err != nil {
			// Log error but don't fail the request
			println("Failed to send notification email:", err.Error())
		}

		// Send auto-reply to sender
		if err := cc.emailService.SendAutoReply(emailData); err != nil {
			// Log error but don't fail the request
			println("Failed to send auto-reply email:", err.Error())
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"message": "Thank you for your message. We will get back to you soon!",
		"id":      message.ID,
	})
}

// GetContactMessages retrieves all contact messages (for admin)
func (cc *ContactController) GetContactMessages(c *gin.Context) {
	var messages []models.ContactMessage

	// Get query parameters
	status := c.Query("status")

	query := cc.db.Order("created_at DESC")

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Limit(50).Find(&messages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve messages",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"messages": messages,
		"count":    len(messages),
	})
}

// UpdateContactMessageStatus updates the status of a contact message
func (cc *ContactController) UpdateContactMessageStatus(c *gin.Context) {
	id := c.Param("id")
	
	var input struct {
		Status string `json:"status" binding:"required,oneof=new read replied"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid status. Must be one of: new, read, replied",
		})
		return
	}

	result := cc.db.Model(&models.ContactMessage{}).
		Where("id = ?", id).
		Update("status", input.Status)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update message status",
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Message not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Status updated successfully",
	})
}

// DeleteContactMessage deletes a contact message
func (cc *ContactController) DeleteContactMessage(c *gin.Context) {
	id := c.Param("id")

	result := cc.db.Delete(&models.ContactMessage{}, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete message",
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Message not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Message deleted successfully",
	})
}

