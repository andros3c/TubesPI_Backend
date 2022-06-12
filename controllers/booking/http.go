package controllers

import (
	"APIRestaurant/bussiness/booking"
	"APIRestaurant/controllers"
	"APIRestaurant/controllers/booking/request"
	"APIRestaurant/controllers/booking/response"
	// "fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookingController struct {
	usecase booking.BookingUseCaseInterface
}

func NewBookingController(uc booking.BookingUseCaseInterface)*BookingController{
	return &BookingController{
		usecase: uc,
	}
}

func (controller *BookingController)CreateNewBooking(c echo.Context)error{
	ctx := c.Request().Context()
	var addBooking request.AddBooking
	err := c.Bind(&addBooking)

	if err != nil{
		return controllers.ErrorResponse(c, http.StatusInternalServerError,"error binding",err)

}

book,err := controller.usecase.CreateNewBooking(*addBooking.ToDomain(),ctx)
if err!=nil{
	return controllers.ErrorResponse(c,http.StatusInternalServerError,"error happen",err)
}
return controllers.SuccesResponse(c,response.FromDomain(book))
}

func (controller *BookingController)GetById(c echo.Context)error{
	ctx := c.Request().Context()
	getId,_ := strconv.Atoi(c.Param("id"))
	book,err := controller.usecase.GetById(getId,ctx)

	if err != nil{
		// if err == gorm.ErrRecordNotFound{}
		// fmt.Println("salah")
		return controllers.ErrorResponse(c, http.StatusInternalServerError,"error happen",err)

}

	return controllers.SuccesResponse(c,book)

}