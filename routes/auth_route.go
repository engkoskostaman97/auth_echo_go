package routes

import (
	

	"project/controllers"
	"project/services"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	
)

func SetAuthRoutes(c *echo.Echo, db *gorm.DB) {
	authController := &controllers.AuthController{Service: &services.UserService{Repo: db}}
	c.POST("/login", authController.Login)
	
}
