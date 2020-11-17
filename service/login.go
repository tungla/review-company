package service

import (
	"review_company/models"
)

type LoginService interface {
	LoginUser(email string, password string) bool
}
type loginInformation struct {
	email    string
	password string
}

func StaticLoginService() LoginService {
	return &loginInformation{
		email:    "bikash.dulal@wesionary.team",
		password: "testing",
	}
}

func (info *loginInformation) LoginUser(email string, password string) bool {
	var user models.User
	if err := models.DB.Where("email = ?", email).First(&user).Error; err == nil {
		hashedPassword := user.Password
		return models.CheckPasswordHash(password, hashedPassword)
	} else {
		return false
	}
}