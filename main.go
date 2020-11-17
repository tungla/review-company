package main

import (
  "github.com/gin-gonic/gin"
  "github.com/swaggo/gin-swagger"
  "github.com/swaggo/files"
  _ "review_company/docs"
  "net/http"

  "review_company/models"
  "review_company/controllers"
  "review_company/service"
  "review_company/middleware"
)

// @title Company Review API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v1
func main() {
  r := gin.Default()

  // Swagger
  url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
  r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

  // Login service
  var loginService service.LoginService = service.StaticLoginService()
  var jwtService service.JWTService = service.JWTAuthService()
  var loginController controllers.LoginController = controllers.LoginHandler(loginService, jwtService)

  models.ConnectDataBase()

  // Company
  r.GET("/companies", middleware.AuthorizeJWT(), controllers.FindCompanies)
  r.POST("/companies", middleware.AuthorizeJWT(), controllers.CreateCompany)
  r.PATCH("/companies/:id", middleware.AuthorizeJWT(), controllers.UpdateCompany)

  // Company Review
  r.GET("/company_reviews/:company_id", middleware.AuthorizeJWT(), controllers.GetCompanyReview)
  r.POST("/company_reviews", middleware.AuthorizeJWT(), controllers.CreateCompanyReview)

  // User
  r.POST("/user", controllers.CreateUser)

  // Login
  r.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Thông tin đăng nhập không hợp lệ!"})
		}
	})

  r.Run()
}