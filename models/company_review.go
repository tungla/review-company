package models

type CompanyReview struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Review  string `json:"review"`
	Rating uint `json:"rating"`
	CompanyID uint `json:"company_id"`
}

type CreateCompanyReviewInput struct {
	Review  string `json:"review"`
	Rating uint `json:"rating"`
	CompanyID uint `valid:"Required" json:"company_id"`
}
