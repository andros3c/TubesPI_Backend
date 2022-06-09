package request

import "APIRestaurant/bussiness/users"

type AddUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
}

func (user *AddUser) ToDomain() *users.DomainUser{
	return &users.DomainUser{
		Email: 		user.Email,
		Password: 	user.Password,
		Name:		user.Name,
		Birthday:  	user.Birthday,	
	}
}