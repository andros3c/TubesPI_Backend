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