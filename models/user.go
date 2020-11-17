package models

import (
  // "github.com/jinzhu/gorm"
)

type User struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  Email  string `json:"email"`
  Password string `json:"password"`
}

type CreateUserInput struct {
  Email  string `json:"email"`
  Password string `json:"password"`
}

type UpdateUserPasswordInput struct {
  Password string `json:"password"`
}