package response

import (
	"APIRestaurant/bussiness/users"
	"time"

	"gorm.io/gorm"
)
type UserResponse struct {
	ID        	uint			`json:"id"`			
	Email	  	string			`json:"email"`
	Name		string			`json:"name"`
	Birthday	string			`json:"birthday"`
	UserRole    string			`json:"user_role"`
	CreatedAt  	time.Time		`json:"created_at"`
	UpdatedAt  	time.Time		`json:"updated_at"`
	DeletedAt 	gorm.DeletedAt  `json:"deleted_at"`
	Token    	string			`json:"token"`


}

func FromDomain(domain users.DomainUser) UserResponse{
	return UserResponse{
		ID: 			domain.ID,
		Email: 			domain.Email,
		Name: 			domain.Name,
		Birthday: 		domain.Birthday,	
		UserRole:  		domain.UserRole,	
		CreatedAt:		domain.CreatedAt,
		UpdatedAt:		domain.UpdatedAt,     
		DeletedAt:		domain.DeletedAt,    
		Token: 			domain.Token, 
	}
}