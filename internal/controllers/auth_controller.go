package controllers

import (
	"net/http"
	"time"

	"manage/internal/models"
	"manage/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	db          *gorm.DB
	authService *services.AuthService
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		db:          db,
		authService: services.NewAuthService(),
	}
}

// LoginRequest represents the login request body
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterRequest represents the register request body
type RegisterRequest struct {
	Username  string `json:"username" binding:"required,min=3,max=50"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Login authenticates an admin and returns a JWT token
func (ac *AuthController) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input",
			"details": err.Error(),
		})
		return
	}

	// Find admin by username
	var admin models.Admin
	if err := ac.db.Where("username = ? AND is_active = ?", req.Username, true).First(&admin).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	// Check password
	if !admin.CheckPassword(req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	// Update last login
	now := time.Now()
	admin.LastLogin = &now
	ac.db.Save(&admin)

	// Generate JWT token
	token, err := ac.authService.GenerateToken(&admin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"admin": gin.H{
			"id":         admin.ID,
			"username":   admin.Username,
			"email":      admin.Email,
			"first_name": admin.FirstName,
			"last_name":  admin.LastName,
		},
	})
}

// Register creates a new admin account
func (ac *AuthController) Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input",
			"details": err.Error(),
		})
		return
	}

	// Check if username already exists
	var existingAdmin models.Admin
	if err := ac.db.Where("username = ?", req.Username).First(&existingAdmin).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Username already exists",
		})
		return
	}

	// Check if email already exists
	if err := ac.db.Where("email = ?", req.Email).First(&existingAdmin).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Email already exists",
		})
		return
	}

	// Create new admin
	admin := models.Admin{
		Username:  req.Username,
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		IsActive:  true,
	}

	// Hash password
	if err := admin.HashPassword(req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// Save to database
	if err := ac.db.Create(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create admin account",
		})
		return
	}

	// Generate JWT token
	token, err := ac.authService.GenerateToken(&admin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Account created but failed to generate token",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Admin account created successfully",
		"token":   token,
		"admin": gin.H{
			"id":         admin.ID,
			"username":   admin.Username,
			"email":      admin.Email,
			"first_name": admin.FirstName,
			"last_name":  admin.LastName,
		},
	})
}

// GetProfile returns the authenticated admin's profile
func (ac *AuthController) GetProfile(c *gin.Context) {
	adminID, exists := c.Get("admin_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	var admin models.Admin
	if err := ac.db.First(&admin, adminID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Admin not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"admin": gin.H{
			"id":          admin.ID,
			"username":    admin.Username,
			"email":       admin.Email,
			"first_name":  admin.FirstName,
			"last_name":   admin.LastName,
			"last_login":  admin.LastLogin,
			"created_at":  admin.CreatedAt,
		},
	})
}

// ChangePassword changes the authenticated admin's password
func (ac *AuthController) ChangePassword(c *gin.Context) {
	adminID, exists := c.Get("admin_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	var req struct {
		CurrentPassword string `json:"current_password" binding:"required"`
		NewPassword     string `json:"new_password" binding:"required,min=8"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input",
			"details": err.Error(),
		})
		return
	}

	var admin models.Admin
	if err := ac.db.First(&admin, adminID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Admin not found",
		})
		return
	}

	// Check current password
	if !admin.CheckPassword(req.CurrentPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Current password is incorrect",
		})
		return
	}

	// Hash and save new password
	if err := admin.HashPassword(req.NewPassword); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash new password",
		})
		return
	}

	if err := ac.db.Save(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Password changed successfully",
	})
}

