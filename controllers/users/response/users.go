package response

import (
	"APIRestaurant/bussiness/users"
	"time"

	"gorm.io/gorm"
)
type UserResponse struct {
	ID        	uint			`json:"id"`			
	Email	  	string			`json:"email"`
	Nama		string			`json:"nama"`
	Birthday	string			`json:"tanggal_lahir"`
	CreatedAt  	time.Time		`json:"created_at"`
	UpdatedAt  	time.Time		`json:"updated_at"`
	DeletedAt 	gorm.DeletedAt  `json:"deleted_at"`
	Token    	string			`json:"token"`


}

func FromDomain(domain users.DomainUser) UserResponse{
	return UserResponse{
		ID: 			domain.ID,
		Email: 			domain.Email,
		Nama: 			domain.Nama,
		Birthday: 		domain.Birthday,		
		CreatedAt:		domain.CreatedAt,
		UpdatedAt:		domain.UpdatedAt,     
		DeletedAt:		domain.DeletedAt,    
		Token: 			domain.Token, 
	}
}