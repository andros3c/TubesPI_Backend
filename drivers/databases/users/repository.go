package users

import (
	"APIRestaurant/bussiness/users"
	"context"
	"errors"

	"gorm.io/gorm"
)
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(gormDb *gorm.DB) users.UserRepoInterface {
	return &UserRepository{
		db:gormDb,
	}
}

func (repo *UserRepository)CreateNewUser(domain users.DomainUser,ctx context.Context)(users.DomainUser,error){
	userDb := FromDomain(domain)

	err := repo.db.Create(&userDb).Error

	if err != nil{
		return users.DomainUser{},err
	}
	return userDb.ToDomain(),nil
}

func (repo *UserRepository)FindEmail(domain users.DomainUser,ctx context.Context)(users.DomainUser,error){
	userDb := FromDomain(domain)

	err := repo.db.Where("email = ?",userDb.Email).First(&userDb).Error
	
	errors.Is(err, gorm.ErrRecordNotFound)
	if err ==  gorm.ErrRecordNotFound{
		return userDb.ToDomain(),gorm.ErrRecordNotFound
	}else if err == nil{
		
		return users.DomainUser{},err
	}else{
		return users.DomainUser{},err
	}
	
}

func (repo *UserRepository) Login(domain users.DomainUser,ctx context.Context)(users.DomainUser,error){
	userDb := FromDomain(domain)
	err:= repo.db.Where("email = ? ",userDb.Email).First(&userDb).Error
	if err != nil{

		if err == gorm.ErrRecordNotFound{
		
			return userDb.ToDomain(),gorm.ErrRecordNotFound
	}
		return users.DomainUser{},err
	}
	return userDb.ToDomain(),nil
}

