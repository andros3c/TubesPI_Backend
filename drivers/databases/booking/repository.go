package booking

import (
	"APIRestaurant/bussiness/booking"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type BookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(gormDb *gorm.DB)booking.BookingRepoInterface{
	return &BookingRepository{
		db:gormDb,
	}
}

func (repo *BookingRepository)CreateNewBooking(domain booking.DomainBooking, ctx context.Context)(booking.DomainBooking,error){
	bookingDb := FromDomain(domain)

	err:= repo.db.Create(&bookingDb).Error

	if err!=nil{
		return booking.DomainBooking{},err
	}
	return bookingDb.ToDomain(),nil
}

func(repo *BookingRepository)GetById(id int , ctx context.Context)(booking.DomainBooking,error){
	book := Booking{}
	
	err := repo.db.Where("id = ?",id).First(&book).Error
	fmt.Println(id)
	if err != nil{
		
		return booking.DomainBooking{},err
	}
	if err == gorm.ErrRecordNotFound{
			return booking.DomainBooking{},gorm.ErrRecordNotFound
		}
	
	return book.ToDomain(),nil

}