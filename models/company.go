package models

import (
  // "github.com/jinzhu/gorm"
)

type Company struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  Name  string `json:"name"`
}

type CreateCompanyInput struct {
  Name  string `valid:"required; MinSize(1)" json:"name"`
}

type UpdateCompanyInput struct {
  Name  string `valid:"required; MinSize(1)" json:"name"`
}