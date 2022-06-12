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
	users.DELETE("/delete/:id",controller.UserController.Delete,middleware.JWTWithConfig(controller.JWTConfig))
	users.PUT("/:id",controller.UserController.Update,middleware.JWTWithConfig(controller.JWTConfig))

	menus := c.Group("/menu")
	menus.POST("/add",controller.MenuController.CreateNewMenu,middleware.JWTWithConfig(controller.JWTConfig))
	menus.GET("/filter/:filter",controller.MenuController.GetAllMenus,middleware.JWTWithConfig(controller.JWTConfig))
	menus.DELETE("/delete/:id",controller.MenuController.DeleteMenu,middleware.JWTWithConfig(controller.JWTConfig))
	menus.GET("/:id",controller.MenuController.FindById,middleware.JWTWithConfig(controller.JWTConfig))

	booking := c.Group("/booking")
	booking.POST("/add", controller.BookingController.CreateNewBooking,middleware.JWTWithConfig(controller.JWTConfig))
	booking.GET("/:id",controller.BookingController.GetById,middleware.JWTWithConfig(controller.JWTConfig))
	booking.POST("/date",controller.BookingController.GetByDate,middleware.JWTWithConfig(controller.JWTConfig))
	booking.GET("/",controller.BookingController.GetAllBookingData,middleware.JWTWithConfig(controller.JWTConfig))
	booking.PUT("/:id",controller.BookingController.UpdateBookingData,middleware.JWTWithConfig(controller.JWTConfig))
}