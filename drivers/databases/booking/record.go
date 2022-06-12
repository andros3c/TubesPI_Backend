package booking

import (
	"APIRestaurant/bussiness/booking"
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	ID            uint
	IDUser        int
	Ordered       string
	TableNumber   int
	Time          string
	Confirmed     bool
	TotalPayment  int
	StatusPayment string
	CreatedAt     time.Time		`gorm:"<-:create"`
	UpdatedAt     time.Time		`gorm:"<-:update"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func FromDomain(domain booking.DomainBooking) Booking{
	return Booking{
		ID: 				domain.ID,
		IDUser: 			domain.IDUser,
		Ordered: 			domain.Ordered,
		TableNumber: 		domain.TableNumber,
		Time: 				domain.Time,
		Confirmed: 			domain.Confirmed,
		TotalPayment: 		domain.TotalPayment,
		StatusPayment: 		domain.StatusPayment,	
		CreatedAt:    		domain.CreatedAt,
		UpdatedAt:   		domain.UpdatedAt,
		DeletedAt:   		domain.DeletedAt,	

	}
}

func (book Booking)ToDomain() booking.DomainBooking{
	return booking.DomainBooking{
		ID: 				book.ID,
		IDUser: 			book.IDUser,
		Ordered: 			book.Ordered,	
		TableNumber: 		book.TableNumber,
		Time: 				book.Time,
		Confirmed: 			book.Confirmed,			
		TotalPayment: 		book.TotalPayment,	
		StatusPayment: 		book.StatusPayment,
		CreatedAt: 			book.CreatedAt,
		UpdatedAt: 			book.UpdatedAt,
		DeletedAt: 			book.DeletedAt,

	}
}