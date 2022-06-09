package menus
import (
	"context"
	"time"

	"gorm.io/gorm"
)

type DomainMenu struct{

	ID 				uint
	MenuName		string
	Category		string
	Price			int
	Desc			string
	Picture			string
	CreatedAt  		time.Time
	UpdatedAt  		time.Time
	DeletedAt 		gorm.DeletedAt `gorm:"index"`

}

type MenuUseCaseInterface interface{
	CreateNewMenu(domain DomainMenu,ctx context.Context)(DomainMenu,error)
	GetAllMenus(ctx context.Context,filter string)([]DomainMenu,error)
}

type MenuRepoInterface interface{
	CreateNewMenu(domain DomainMenu,ctx context.Context)(DomainMenu,error)
	GetAllMenus(ctx context.Context, filter string)([]DomainMenu,error)
	
}