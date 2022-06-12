package menus

import (
	"APIRestaurant/bussiness/menus"
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	ID        		uint
	MenuName  		string
	Category  		string
	Price     		int
	Desc      		string
	Picture   		string
	CreatedAt 		time.Time `gorm:"<-:create"`
	UpdatedAt 		time.Time	`gorm:"<-:update"`
	DeletedAt 		gorm.DeletedAt `gorm:"index"`
}

func (menu Menu) ToDomain() menus.DomainMenu{
	return menus.DomainMenu{
		ID: menu.ID,
		MenuName: 		menu.MenuName,
		Category: 		menu.Category,
		Price: 			menu.Price,
		Desc: 			menu.Desc,
		Picture: 		menu.Picture,
		CreatedAt: 		menu.CreatedAt,
		UpdatedAt: 		menu.UpdatedAt,
		DeletedAt: 		menu.DeletedAt,
	}
}

func FromDomain(domain menus.DomainMenu) Menu{
	return Menu{
		ID: domain.ID,
		MenuName: 		domain.MenuName,
		Category:		domain.Category,
		Price: 			domain.Price,
		Desc: 			domain.Desc,
		Picture: 		domain.Picture,
		CreatedAt:    	domain.CreatedAt,
		UpdatedAt:   	domain.UpdatedAt,
		DeletedAt:   	domain.DeletedAt,	
	}
}