package controllers

import (
	"github.com/gin-gonic/gin"
	"review_company/models"
	"net/http"
	"github.com/astaxie/beego/validation"
)

// GetCompanyReview godoc
// @Summary Get all reviews of a company
// @Description Get all reviews of a company by company id
// @Tags Company Review
// @Produce  json
// @Param company_id path int true "Company ID"
// @Param token header string true "Token"
// @Success 200 {object} models.CompanyReview
// @Router /company_reviews/{company_id} [get]
func GetCompanyReview(c *gin.Context) {
	var companyReviews []models.CompanyReview
	models.DB.Where("company_id = ?", c.Param("company_id")).Find(&companyReviews)

	c.JSON(http.StatusOK, gin.H{"data": companyReviews})
}

// CreateCompanyReview godoc
// @Summary Craete review for a company
// @Description Craete review for a company
// @Tags Company Review
// @Accept  json
// @Produce  json
// @Param company_id body int true "Company ID"
// @Param rating body int true "Rating"
// @Param review body string true "Review"
// @Param token header string true "Token"
// @Success 200 {object} models.CompanyReview
// @Router /company_reviews [post]
func CreateCompanyReview(c *gin.Context) {
	var input models.CreateCompanyReviewInput
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
		// Create company review
		review := models.CompanyReview{Review: input.Review, Rating: input.Rating, CompanyID: input.CompanyID}
		models.DB.Create(&review)
		c.JSON(http.StatusOK, gin.H{"data": review})
	}

}