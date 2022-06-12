package routes

import (
	userController "APIRestaurant/controllers/users"
	menuController "APIRestaurant/controllers/menus"
	bookingController  "APIRestaurant/controllers/booking"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
type RouteControllerList struct{
	UserController userController.UserController
	MenuController menuController.MenuController 
	BookingController bookingController.BookingController 
	JWTConfig middleware.JWTConfig
}

func (controller RouteControllerList)RouteRegister (c *echo.Echo){
	users := c.Group("/user")
	users.POST("/add",controller.UserController.CreateNewUser)
	users.POST("/login",controller.UserController.Login)

	menus := c.Group("/menu")
	menus.POST("/add",controller.MenuController.CreateNewMenu,middleware.JWTWithConfig(controller.JWTConfig))
	menus.GET("/:filter",controller.MenuController.GetAllMenus,middleware.JWTWithConfig(controller.JWTConfig))
	menus.DELETE("/delete/:id",controller.MenuController.DeleteMenu,middleware.JWTWithConfig(controller.JWTConfig))

	booking := c.Group("/booking")
	booking.POST("/add", controller.BookingController.CreateNewBooking,middleware.JWTWithConfig(controller.JWTConfig))
	booking.GET("/:id",controller.BookingController.GetById,middleware.JWTWithConfig(controller.JWTConfig))
}