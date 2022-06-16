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

func (repo *UserRepository)Delete(id int,ctx context.Context)(users.DomainUser,error){
	user := User{}

	err:=repo.db.Delete(&user,id).Error

	if err != nil{
		return users.DomainUser{},err
	}
	return user.ToDomain(),nil
}




func (repo *UserRepository) FindById(userId int,ctx context.Context)(users.DomainUser,error){
	user := User{}

	err := repo.db.Where("id = ?",userId).First(&user).Error
	if err != nil{
		return users.DomainUser{},err
	}
	if err == gorm.ErrRecordNotFound{
		return users.DomainUser{},gorm.ErrRecordNotFound
	}
	return user.ToDomain(),nil
}

func (repo *UserRepository) Update(userDomain *users.DomainUser, ctx context.Context) (users.DomainUser, error) {
	rec := FromDomain(*userDomain)

	result := repo.db.Save(&rec)
	if result.Error != nil {
		return users.DomainUser{}, result.Error
	}

	record, err := repo.FindById(int(rec.ID), ctx)
	if err != nil {
		return users.DomainUser{}, err
	}

	return record, nil
}
