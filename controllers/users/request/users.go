package request

import "APIRestaurant/bussiness/users"

type AddUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Nama     string `json:"nama"`
	Birthday string `json:"tanggal_lahir"`
}

func (user *AddUser) ToDomain() *users.DomainUser{
	return &users.DomainUser{
		Email: 		user.Email,
		Password: 	user.Password,
		Nama:		user.Nama,
		Birthday:  	user.Birthday,	
	}
}