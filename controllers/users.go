package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    "net/http"
    "log"
    "fmt"

    "crud_go_app/models"
    "crud_go_app/auth"
)

type CreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Signup godoc
// @Summary Create a new User
// @Description post request example
// @Accept json
// @Produce plain
// @Param message body models.User true "User Info"
// @Success 200 {string} string "success"
// @Failure 500 {string} string "fail"
// @Router /public/login/ [post]
func Signup(c *gin.Context) {
  // Signup creates a user in db
	var user models.User
  
  fmt.Println(c.Request)

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)

		c.JSON(400, gin.H{
			"msg": "invalid json",
		})
		c.Abort()

		return
	}

	err = user.HashPassword(user.Password)
	if err != nil {
		log.Println(err.Error())

		c.JSON(500, gin.H{
			"msg": "error hashing password",
		})
		c.Abort()

		return
	}

	err = user.CreateUserRecord()
	if err != nil {
		log.Println(err)

		c.JSON(500, gin.H{
			"msg": "error creating user",
		})
		c.Abort()

		return
	}

	c.JSON(200, user)
}


// LoginPayload login body
type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse token response
type LoginResponse struct {
	Token string `json:"token"`
}


// Login logs users in
func Login(c *gin.Context) {
	var payload LoginPayload
	var user models.User

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	}

	result := models.DB.Where("email = ?", payload.Email).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(401, gin.H{
			"msg": "invalid user credentials",
		})
		c.Abort()
		return
	}

	err = user.CheckPassword(payload.Password)
	if err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"msg": "invalid user credentials",
		})
		c.Abort()
		return
	}

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "verysecretkey",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(user.Email)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"msg": "error signing token",
		})
		c.Abort()
		return
	}

	tokenResponse := LoginResponse{
		Token: signedToken,
	}

	c.JSON(http.StatusOK, tokenResponse)

	return
}
