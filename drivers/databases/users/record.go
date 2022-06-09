package users

import (
	"APIRestaurant/bussiness/users"
	"time"

	"gorm.io/gorm"
)
type User struct {
	ID        	uint
	Email	  	string
	Password    string
	Nama		string
	Birthday	string
	CreatedAt  	time.Time		`gorm:"<-:create"`		
	UpdatedAt  	time.Time
	DeletedAt 	gorm.DeletedAt  `gorm:"index"`
	

}

func (User User) ToDomain() users.DomainUser{
	return users.DomainUser{
		ID: 			User.ID,
		Email: 			User.Email,	
		Password: 		User.Password,
		Nama: 			User.Nama,
		Birthday: 		User.Birthday,	
		CreatedAt: 		User.CreatedAt,
		UpdatedAt: 		User.UpdatedAt ,	
		DeletedAt: 		User.DeletedAt,
	}
}

func FromDomain(domain users.DomainUser) User{
	return User{
		ID: 			domain.ID,
		Email:			domain.Email,
		Password: 		domain.Password,
		Nama:			domain.Nama,
		Birthday: 		domain.Birthday,
		CreatedAt:    	domain.CreatedAt,
		UpdatedAt:   	domain.UpdatedAt,
		DeletedAt:   	domain.DeletedAt,	
		
	}
}