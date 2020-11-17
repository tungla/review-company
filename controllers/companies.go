package controllers

import (
  "github.com/astaxie/beego/validation"
  "github.com/gin-gonic/gin"
  "review_company/models"
  "net/http"
)

// FindCompanies godoc
// @Summary Get company list
// @Description Get all companies
// @Tags Company
// @Produce  json
// @Param token header string true "Token"
// @Success 200 {object} models.Company
// @Router /companies [get]
func FindCompanies(c *gin.Context) {
  var companies []models.Company
  models.DB.Find(&companies)

  c.JSON(http.StatusOK, gin.H{"data": companies})
}

// CreateCompany godoc
// @Summary Create a company
// @Description Create a company
// @Tags Company
// @Accept  json
// @Produce  json
// @Param name body string true "Name"
// @Param token header string true "Token"
// @Success 200 {object} models.Company
// @Router /companies [post]
func CreateCompany(c *gin.Context) {
  var input models.CreateCompanyInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Validate input
  valid := validation.Validation{}
  b, err := valid.Valid(&input)
  if err != nil {
    // handle error
  }
  if !b {
    msg := ""
    for _, err := range valid.Errors {
      msg = msg + err.Message
    }
    c.JSON(http.StatusBadRequest, gin.H{"err_msg": msg})
  } else {
    // Create company
    company := models.Company{Name: input.Name}
    models.DB.Create(&company)
    c.JSON(http.StatusOK, gin.H{"data": company})
  }
}

// UpdateCompany godoc
// @Summary Update name of a company
// @Description Update name of a company
// @Tags Company
// @Accept  json
// @Produce  json
// @Param company_id path int true "Company ID"
// @Param name body string true "Name"
// @Param token header string true "Token"
// @Success 200 {object} models.Company
// @Router /company_reviews/{company_id} [patch]
func UpdateCompany(c *gin.Context) {
  // Get model if exist
  var company models.Company
  if err := models.DB.Where("id = ?", c.Param("id")).First(&company).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Không tìm thấy công ty nào với id này!"})
    return
  }

  // Validate json input
  var input models.UpdateCompanyInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Validate input
  valid := validation.Validation{}
  b, err := valid.Valid(&input)
  if err != nil {
    // handle error
  }
  if !b {
    msg := ""
    for _, err := range valid.Errors {
      msg = msg + err.Message
    }
    c.JSON(http.StatusBadRequest, gin.H{"err_msg": msg})
  } else {
    models.DB.Model(&company).Updates(input)
    c.JSON(http.StatusOK, gin.H{"data": company})
  }
}