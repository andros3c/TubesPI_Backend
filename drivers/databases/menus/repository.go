package menus

import (
	"APIRestaurant/bussiness/menus"
	"context"

	"gorm.io/gorm"
)

type MenuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(gormDb *gorm.DB)menus.MenuRepoInterface{
	return &MenuRepository{
		db:gormDb,
	}
}

func (repo *MenuRepository)CreateNewMenu(domain menus.DomainMenu, ctx context.Context)(menus.DomainMenu,error){
	menuDb := FromDomain(domain)

	err:= repo.db.Create(&menuDb).Error

	if err != nil{
		return menus.DomainMenu{},err
	}
	return menuDb.ToDomain(),nil
}

func (repo *MenuRepository)GetAllMenus(ctx context.Context, filter string)([]menus.DomainMenu,error){
	menuDb := []Menu{}
	MenusDomain := []menus.DomainMenu{}

	err := repo.db.Error

	if filter == "newest"{
		err = repo.db.Order("created_at DESC").Find(&menuDb).Error
		if err != nil {
			return []menus.DomainMenu{}, err
		}
	}else if filter == "all"{
		err = repo.db.Find(&menuDb).Error
		if err != nil {
			return []menus.DomainMenu{}, err
		}
	
	}else if filter != ""{
		err = repo.db.Where("category = ?",filter).Find(&menuDb).Error
		if err != nil {
			return []menus.DomainMenu{}, err
		}
	}
	
	for _,value := range menuDb{
		MenusDomain = append(MenusDomain, value.ToDomain())
	}
	return MenusDomain,nil
}

func (repo *MenuRepository)DeleteMenu(ctx context.Context,id int)(menus.DomainMenu,error){
	menu:= Menu{}

		
	res := repo.db.Delete(&menu,id)
	if res.Error != nil{
		return menus.DomainMenu{},res.Error
	}
	return menu.ToDomain(),nil

	}
	
func(repo *MenuRepository)GetById(id int , ctx context.Context)(menus.DomainMenu,error){
	menu := Menu{}

	err := repo.db.Where("id = ?",id).First(&menu).Error

	if err != nil{
		return menus.DomainMenu{},err
	}
	if err == gorm.ErrRecordNotFound{
			return  menus.DomainMenu{},gorm.ErrRecordNotFound
		}
	
	return menu.ToDomain(),nil
}