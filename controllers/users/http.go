package controllers

import (
	"APIRestaurant/bussiness"
	"APIRestaurant/bussiness/users"
	"APIRestaurant/controllers"
	"APIRestaurant/controllers/users/request"
	"APIRestaurant/controllers/users/response"
	"fmt"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)
type UserController struct {
	usecase users.UserUseCaseInterface
}

func NewUserController(uc users.UserUseCaseInterface)*UserController{
	return &UserController{
		usecase: uc,
	}
}

func (controller *UserController) CreateNewUser(c echo.Context)error{
	ctx := c.Request().Context()
	var addUser request.AddUser
	err := c.Bind(&addUser)

	if err != nil{
			return controllers.ErrorResponse(c, http.StatusInternalServerError,"error binding",err)

	}
	prod,err := controller.usecase.CreateNewUser(*addUser.ToDomain(),ctx)
	if err != nil{

		if err == bussiness.ErrDuplicateEmail{
			return controllers.ErrorResponse(c, http.StatusConflict,"Dupicate data",err)
		}else{
			return controllers.ErrorResponse(c,http.StatusInternalServerError,"Error Happen",err)
		}
	}
	return controllers.SuccesResponse(c,response.FromDomain(prod))
}

func (controller *UserController) Login(c echo.Context) error{
	ctx := c.Request().Context()
	var userLogin request.AddUser
	err := c.Bind(&userLogin)

	if err != nil{
		return controllers.ErrorResponse(c,http.StatusInternalServerError,"Error Happen",err)
	}

	user,err := controller.usecase.Login(*userLogin.ToDomain(),ctx)

	if err != nil{
		
		return controllers.ErrorResponse(c,http.StatusInternalServerError,"Error Happen",err)
		
	}
	return controllers.SuccesResponse(c,response.FromDomain(user))
}

func (controller *UserController)Delete(c echo.Context)error{
	ctx := c.Request().Context()
	userId,_:= strconv.Atoi(c.Param("id"))

	deleteUser,err := controller.usecase.Delete(userId,ctx)

	if err != nil{
		return controllers.ErrorResponse(c,http.StatusInternalServerError,err.Error(),err)
	}
	return controllers.SuccesResponse(c,response.FromDomain(deleteUser))
}

func (ctrl *UserController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	idstr := c.Param("id")

	// if strings.TrimSpace(idstr) == "" {
	// 	return controllers.ErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	// }

	id, _ := strconv.Atoi(idstr)

	req := request.AddUser{}
	if err := c.Bind(&req)
	 err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest,"Error Binding", err)
	}

	domainReq := req.ToDomain()
	domainReq.ID = uint(id)
	fmt.Println("domainReq.ID")
	resp, err := ctrl.usecase.Update(domainReq, ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError,"Error Happen", err)
	}

	return controllers.SuccesResponse(c, response.FromDomain(*resp))
}