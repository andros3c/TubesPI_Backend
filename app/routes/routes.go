package routes

import (
	userController "APIRestaurant/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
type RouteControllerList struct{
	UserController userController.UserController
	JWTConfig middleware.JWTConfig
}

func (controller RouteControllerList)RouteRegister (c *echo.Echo){
	users := c.Group("/user")

	users.POST("/add",controller.UserController.CreateNewUser)
	users.POST("/login",controller.UserController.Login)
}