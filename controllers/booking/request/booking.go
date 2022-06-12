package request

import "APIRestaurant/bussiness/booking"

type AddBooking struct {
	IDUser        int    `json:"user_id"`
	Ordered       string `json:"order"`
	TableNumber   int    `json:"table_number"`
	Time          string `json:"time"`
	Confirmed     bool   `json:"confirmed_status"`
	TotalPayment  int    `json:"payment_total"`
	StatusPayment string `json:"payment_status"`
}

func (book *AddBooking) ToDomain() *booking.DomainBooking{
	return &booking.DomainBooking{
		IDUser:  		book.IDUser,
		Ordered: 		book.Ordered,
		TableNumber: 	book.TableNumber,
		Time: 			book.Time,
		Confirmed: 		book.Confirmed,	
		TotalPayment: 	book.TotalPayment,
		StatusPayment: 	book.StatusPayment,	
	}
}