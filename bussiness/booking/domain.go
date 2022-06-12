package booking

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type DomainBooking struct {
	ID            uint
	IDUser        int
	Ordered       string
	TableNumber   int
	Date          string
	Time          string
	Confirmed     bool
	TotalPayment  int
	StatusPayment string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type BookingUseCaseInterface interface{
	CreateNewBooking(domain DomainBooking, ctx context.Context)(DomainBooking,error)
	GetById(id int , ctx context.Context)(DomainBooking,error)
	GetByDate(domain  DomainBooking , ctx context.Context)([]DomainBooking,error)
	GetAllBookingData(domain DomainBooking,ctx context.Context)([]DomainBooking,error)
}

type BookingRepoInterface interface{
	CreateNewBooking(domain DomainBooking, ctx context.Context)(DomainBooking,error)
	GetById(id int , ctx context.Context)(DomainBooking,error)
	GetByDate(domain  DomainBooking , ctx context.Context)([]DomainBooking,error)
	GetAllBookingData(domain DomainBooking,ctx context.Context)([]DomainBooking,error)

}