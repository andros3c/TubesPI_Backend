package booking

import (
	"APIRestaurant/bussiness/booking"
	"context"
	

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
	// fmt.Println(id)
	if err != nil{
		
		return booking.DomainBooking{},err
	}
	if err == gorm.ErrRecordNotFound{
			return booking.DomainBooking{},gorm.ErrRecordNotFound
		}
	
	return book.ToDomain(),nil

}

func(repo *BookingRepository)GetByDate(domain  booking.DomainBooking , ctx context.Context)([]booking.DomainBooking,error){
	bookByDate := []Booking{}
	BookingDomain := []booking.DomainBooking{}
	
	err := repo.db.Where("date = ?",domain.Date).Find(&bookByDate).Error
	
	if err != nil{
		
		return []booking.DomainBooking{},err
	}
	if err == gorm.ErrRecordNotFound{
			return []booking.DomainBooking{},gorm.ErrRecordNotFound
		}
	
		for _,value := range bookByDate{
			BookingDomain = append(BookingDomain, value.ToDomain())
		}
	return BookingDomain,nil
}

func(repo *BookingRepository)GetAllBookingData(domain booking.DomainBooking,ctx context.Context)([]booking.DomainBooking,error){
	bookingDb := []Booking{}
	bookingDomain := []booking.DomainBooking{}

	err := repo.db.Find(&bookingDb).Error
	if err != nil{
		return []booking.DomainBooking{},err
	}
	return bookingDomain,nil
}