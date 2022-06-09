package request

import "APIRestaurant/bussiness/menus"

type AddMenu struct {
	MenuName string `json:"menu_name"`
	Category string `json:"menu_category"`
	Price    int    `json:"menu_price"`
	Desc     string `json:"menu_description"`
	Picture  string `json:"menu_picture"`
}

func (menu *AddMenu) ToDomain() *menus.DomainMenu{
	return &menus.DomainMenu{
		MenuName: 		menu.MenuName,
		Category:		menu.Category,
		Price: 			menu.Price,
		Desc: 			menu.Desc,
		Picture: 		menu.Picture,
	}
}