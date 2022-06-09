package response

import (
	"APIRestaurant/bussiness/menus"
	"time"

	"gorm.io/gorm"
)

type MenuResponse struct {
	ID 				uint
	MenuName		string			`json:"menu_name"`
	Category		string			`json:"menu_category"`
	Price			int				`json:"menu_price"`
	Desc			string			`json:"menu_description"`
	Picture			string			`json:"menu_picture"`
	CreatedAt  		time.Time		`json:"created_at"`
	UpdatedAt  		time.Time		`json:"updated_at"`
	DeletedAt 		gorm.DeletedAt 	`json:"deleted_at"`
}

func FromDomain(domain menus.DomainMenu) MenuResponse{
	return MenuResponse{
		ID:				domain.ID,
		MenuName: 		domain.MenuName,
		Category: 		domain.Category,
		Price: 			domain.Price,
		Desc: 			domain.Desc,
		Picture: 		domain.Picture,
		CreatedAt: 		domain.CreatedAt,
		UpdatedAt: 		domain.UpdatedAt,
		DeletedAt:		domain.DeletedAt,
	}
}