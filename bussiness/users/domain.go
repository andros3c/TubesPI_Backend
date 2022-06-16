package users

import (
	"context"
	"time"

	"gorm.io/gorm"
)
type DomainUser struct {
	ID        	uint
	Email	  	string
	Password    string
	Name		string
	Birthday	string
	UserRole    string
	CreatedAt  	time.Time
	UpdatedAt  	time.Time
	DeletedAt 	gorm.DeletedAt `gorm:"index"`
	Token    	string
}

type UserUseCaseInterface interface{
	CreateNewUser(domain DomainUser,ctx context.Context)(DomainUser,error)
	Login(domain DomainUser,ctx context.Context)(DomainUser,error)
	Delete(id int,ctx context.Context)(DomainUser,error)
	Update(userDomain *DomainUser, ctx context.Context, ) (*DomainUser, error)
}
type UserRepoInterface interface{
	CreateNewUser(domain DomainUser,ctx context.Context)(DomainUser,error)
	FindEmail(domain DomainUser,ctx context.Context)(DomainUser,error)
	Login(domain DomainUser,ctx context.Context)(DomainUser,error)
	Delete(id int,ctx context.Context)(DomainUser,error)
	Update(userDomain *DomainUser, ctx context.Context) (DomainUser, error)
	FindById(userId int,ctx context.Context)(DomainUser,error)
}