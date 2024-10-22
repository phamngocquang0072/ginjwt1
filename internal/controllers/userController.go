package controllers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/quangnpham/ginjwt1/initializers"
	"github.com/quangnpham/ginjwt1/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Response struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   string `json:"data"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func init() {
	// Load .env file
	initializers.LoadEnv()
	DB = initializers.ConnectDB()
}

func GetUser(c *gin.Context) {
	users := []models.User{}
	DB.Find(&users)

	c.JSON(200, gin.H{
		"users": users,
	})
}

// @Summary      Create a user
// @Description  Create a single user or multiple users
// @Tags         users
// @Accept       json
// @Produce      application/json
// @Param        user  body  UserResponse  true  "User details"
// @Success      201  {object}  map[string]interface{}
// @Failure      400  {object} map[string]interface{}
// @Router       /user [post]
func CreateUser(c *gin.Context) {
	// Initialize a slice for Users

	// Try to bind as a single User first
	singleUser := models.User{}
	if err := c.ShouldBindJSON(&singleUser); err == nil {
		hash, err := bcrypt.GenerateFromPassword([]byte(singleUser.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
			return
		}

		singleUser.Password = string(hash)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Perform bulk creation of Users
	result := DB.Create(&singleUser)
	if result.Error != nil {
		log.Println("Failed to create Users:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create Users"})
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, gin.H{"message": "Users added successfully"})
}

func DeleteUser(c *gin.Context) {
	var user models.User
	if err := DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func Login(c *gin.Context) {
	var body struct {
		Username string `json:"username"` // Capitalized
		Password string `json:"password"` // Capitalized
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := DB.Where("username = ?", body.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid username"})
		return
	}

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		log.Println("Invalid password:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	// Generate token if password is correct
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create token"})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"user":    user,
	})
}
