package menus

import (
	"APIRestaurant/bussiness"
	"context"
	"time"
)

type MenuUseCase struct {
	repo MenuRepoInterface
	ctx  time.Duration
}

func NewMenuUseCase(MenuRepo  MenuRepoInterface, contextTimeout time.Duration)MenuUseCaseInterface{
	return &MenuUseCase{
		repo : MenuRepo,
		ctx  : contextTimeout,
	}
}


func (usecase *MenuUseCase)CreateNewMenu(domain DomainMenu,ctx context.Context)(DomainMenu,error){
	if domain.MenuName == ""{
		return DomainMenu{},bussiness.ErrMenuNameEmpty
	}
	if domain.Price == 0 {
		return DomainMenu{},bussiness.ErrMenuPriceEmpty
	}
	if domain.Category== ""{
		return DomainMenu{},bussiness.ErrCategoryEmpty
	}
	if domain.Picture== ""{
		return DomainMenu{},bussiness.ErrPictureEmpty
}

menu, err := usecase.repo.CreateNewMenu(domain,ctx)
if err != nil{
	return DomainMenu{},err
}
return menu,nil

}

func (usecase *MenuUseCase) GetAllMenus(ctx context.Context,filter string)([]DomainMenu,error){
	result,err := usecase.repo.GetAllMenus(ctx,filter)
	if err != nil{
		return []DomainMenu{},err
	}
	return result,nil
}

func (usecase *MenuUseCase) DeleteMenu(ctx context.Context,id int)(DomainMenu,error){
	menu,err := usecase.repo.DeleteMenu(ctx,id)
	if err != nil{
		return DomainMenu{},err
	}
	return menu,nil
}