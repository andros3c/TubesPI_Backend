package users

import (
	"APIRestaurant/bussiness"
	"APIRestaurant/drivers/helpers/encrypt"
	"context"
	"time"
	_middleware "APIRestaurant/app/middleware"
	
	"gorm.io/gorm"
)

type UserUseCase struct {
	repo UserRepoInterface
	ctx  time.Duration
	jwt	 *_middleware.ConfigJWT
}

func NewUserUseCase(UserRepo UserRepoInterface, contextTimeout time.Duration,configJWT *_middleware.ConfigJWT)UserUseCaseInterface{
	return &UserUseCase{
		repo: UserRepo,
		ctx: contextTimeout,
		jwt: configJWT,
	}
}

func (usecase *UserUseCase)CreateNewUser(domain DomainUser,ctx context.Context)(DomainUser,error){
	if domain.Email == ""{
		return DomainUser{},bussiness.ErrEmailEmpty
	}
	if domain.Password == ""{
		return DomainUser{},bussiness.ErrPassEmpty
	}
	if domain.Nama == ""{
		return DomainUser{},bussiness.ErrNameEmpty
	}

	existedEmail,err := usecase.repo.FindEmail(domain,ctx)
	if err == gorm.ErrRecordNotFound{
		domain.Password,err = encrypt.Hash(domain.Password)
		if err != nil{
			return DomainUser{},bussiness.ErrInternalServer
		}
		user,err := usecase.repo.CreateNewUser(domain,ctx)
		if err != nil{
			return DomainUser{},err
		}
		return user,nil
	}else{
		return existedEmail, bussiness.ErrDuplicateEmail
	}
	
	
}
func (usecase *UserUseCase)Login(domain DomainUser,ctx context.Context)(DomainUser,error){
	if domain.Email == ""{
		return DomainUser{}, bussiness.ErrEmailEmpty
	}
	if domain.Password == ""{
		return DomainUser{},bussiness.ErrPassEmpty
	}

	user,err := usecase.repo.Login(domain,ctx)
	if err!=nil{

		if err == gorm.ErrRecordNotFound{
			return DomainUser{},bussiness.ErrAccNotFound
		}

		return DomainUser{},err

	}
	if !encrypt.ValidateHash(domain.Password,user.Password){
		return DomainUser{},bussiness.ErrWrongPass
	}
	user.Token = usecase.jwt.GenererateToken(user.ID)
	return user,nil

}