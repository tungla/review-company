package controllers

import (
  "github.com/gin-gonic/gin"
  "review_company/models"
  "net/http"
)

// POST /companies
// Create new user
func CreateUser(c *gin.Context) {
  var input models.CreateUserInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Validate unique email
  var user models.User
  if err := models.DB.Where("email = ?", input.Email).First(&user).Error; err == nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Email này đã được đăng ký, vui lòng sử dụng email khác!"})
    return
  }

  // Create user
  pass, err := models.HashPasswordBcrypt(input.Password)
  if err == nil {
    user := models.User{Email: input.Email, Password: pass}
    models.DB.Create(&user)
    c.JSON(http.StatusOK, gin.H{"data": user})
  } else {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  }
}