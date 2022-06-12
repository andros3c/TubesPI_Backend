package booking

import (
	
	"context"
	"time"
)

type BookingUseCase struct {
	repo BookingRepoInterface
	ctx  time.Duration
}

func NewBookingUseCase(BookingRepo BookingRepoInterface,contextTimeout time.Duration)BookingUseCaseInterface{
	return &BookingUseCase{
		repo : BookingRepo,
		ctx  : contextTimeout,
	}
}

func (usecase *BookingUseCase)CreateNewBooking(domain DomainBooking,ctx context.Context)(DomainBooking,error){
	booking,err := usecase.repo.CreateNewBooking(domain,ctx)

	if err != nil{
		return DomainBooking{},err
	}
	return booking,nil
}

func (usecase *BookingUseCase)GetById(id int , ctx context.Context)(DomainBooking,error){
		booking,err := usecase.repo.GetById(id,ctx)

		if err != nil{
			return DomainBooking{},err
		}
		return booking,nil
}

func (usecase *BookingUseCase)GetByDate(domain  DomainBooking, ctx context.Context)([]DomainBooking,error){
	booking,err := usecase.repo.GetByDate(domain,ctx)

		if err != nil{
			return []DomainBooking{},err
		}
		return booking,nil
}

func (usecase *BookingUseCase)GetAllBookingData(domain DomainBooking,ctx context.Context)([]DomainBooking,error){
	booking,err := usecase.repo.GetAllBookingData(domain,ctx) 
	if err != nil{
		return []DomainBooking{},ctx.Err()

	}
	return booking,nil
}