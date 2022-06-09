package controllers

import (
	"APIRestaurant/bussiness/menus"
	"APIRestaurant/controllers"
	"APIRestaurant/controllers/menus/request"
	"APIRestaurant/controllers/menus/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MenuController struct {
	usecase menus.MenuUseCaseInterface
}

func NewMenuController(uc menus.MenuUseCaseInterface)*MenuController{
	return &MenuController{
		usecase: uc,
	}
}

func (controller *MenuController) CreateNewMenu(c echo.Context)error{
	ctx := c.Request().Context()
	var addMenu request.AddMenu
	err := c.Bind(&addMenu)

	if err != nil{
		return controllers.ErrorResponse(c, http.StatusInternalServerError,"error binding",err)

}

	create,err := controller.usecase.CreateNewMenu(*addMenu.ToDomain(),ctx)
	if err != nil{
		return controllers.ErrorResponse(c,http.StatusInternalServerError,"Error Happen",err)
	}
	return controllers.SuccesResponse(c,response.FromDomain(create))
}

func (controller *MenuController) GetAllMenus(c echo.Context)error{
	ctx := c.Request().Context()
	filter := c.Param("filter")
	menu,err := controller.usecase.GetAllMenus(ctx,filter)

	if err != nil{
		return controllers.ErrorResponse(c,http.StatusInternalServerError,"Error Happen",err)
	}

	respon := []response.MenuResponse{}

	for _, values := range menu{
		respon = append(respon,response.FromDomain(values))
	}
return controllers.SuccesResponse(c,respon)
}