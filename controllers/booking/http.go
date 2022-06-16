package controllers

import (
	"APIRestaurant/bussiness/booking"
	"APIRestaurant/controllers"
	"APIRestaurant/controllers/booking/request"
	"APIRestaurant/controllers/booking/response"
	"fmt"

	// "fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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

func (controller *BookingController)GetByDate(c echo.Context)error{
	ctx := c.Request().Context()
	var getDate request.AddBooking
	err := c.Bind(&getDate)

	if err != nil{
		return controllers.ErrorResponse(c, http.StatusInternalServerError,"error binding",err)

}

	
	fmt.Println(getDate)

	book,err := controller.usecase.GetByDate(*getDate.ToDomain(),ctx) 

	if err != nil {
		return controllers.ErrorResponse(c,http.StatusInternalServerError,"Error Happen",err)

	}

	respons := []response.BookingResponse{}
	
	for _, values := range book{
		respons = append(respons, response.FromDomain(values))
	}
	return controllers.SuccesResponse(c,respons)
}

func (controller *BookingController)GetAllBookingData(c echo.Context)error{
	ctx := c.Request().Context()
	// var bookDomain request.AddBooking

	// err := c.Bind(&bookDomain)

// 	if err != nil{
// 		return controllers.ErrorResponse(c, http.StatusInternalServerError,"error binding",err)

// }
	book,err := controller.usecase.GetAllBookingData(ctx )
	if err != nil{
		return controllers.ErrorResponse(c,http.StatusInternalServerError,"Error Happen",err)
	}
	respon := []response.BookingResponse{}

	for _,values := range book{
		respon = append(respon, response.FromDomain(values))
	}

	return controllers.SuccesResponse(c,respon)
}

func (controller *BookingController)UpdateBookingData(c echo.Context)error{
	ctx := c.Request().Context()
	bookId, _ := strconv.Atoi(c.Param("id"))
	book  := request.AddBooking{}

	c.Bind(&book)
domainReq := book.ToDomain()
	booking,err := controller.usecase.UpdateBookingData(bookId,*domainReq,ctx)
	if err != nil{
		return controllers.ErrorResponse(c,http.StatusInternalServerError,err.Error(),err)
	}
	return controllers.SuccesResponse(c,response.FromDomain(booking))
}

func (controller *BookingController)DeleteBookingData(c echo.Context)error{
	ctx := c.Request().Context()
	BookId, _ := strconv.Atoi(c.Param("id"))
	book,err := controller.usecase.DeleteBookingData(BookId,ctx)
	if err != nil{
		if err == gorm.ErrRecordNotFound{}
		return controllers.ErrorResponse(c,http.StatusInternalServerError,"Error Happen",err)
	}
	return controllers.SuccesResponse(c,book)
}