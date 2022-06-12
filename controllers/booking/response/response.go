package response

import (
	"APIRestaurant/bussiness/booking"
	"time"

	"gorm.io/gorm"
)

type BookingResponse struct {
	ID            uint   			`json:"booking_id"`
	IDUser        int    			`json:"user_id"`
	Ordered       string 			`json:"order"`
	TableNumber   int    			`json:"table_number"`
	Date 		  string			`json:"date"`
	Time          string 			`json:"time"`
	Confirmed     bool   			`json:"confirmed_status"`
	TotalPayment  int    			`json:"payment_total"`
	StatusPayment string 			`json:"payment_status"`
	CreatedAt     time.Time			`json:"created_at"`
	UpdatedAt     time.Time			`json:"updated_at"`
	DeletedAt     gorm.DeletedAt	`json:"deleted_at"`
}

func FromDomain(domain booking.DomainBooking) BookingResponse{
	return BookingResponse{
		ID: 				domain.ID,	
		IDUser: 			domain.IDUser,
		Ordered:  			domain.Ordered,
		TableNumber:  		domain.TableNumber,
		Date: 				domain.Date,
		Time:  				domain.Time,
		Confirmed:  		domain.Confirmed,
		TotalPayment:  		domain.TotalPayment,
		StatusPayment:  	domain.StatusPayment,
		CreatedAt: 			domain.CreatedAt,
		UpdatedAt: 			domain.UpdatedAt,
		DeletedAt:			domain.DeletedAt,
	}
}