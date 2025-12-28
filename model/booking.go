package model

type Booking struct {
	Qty       int    `json:"qty"`
	CreatedAt string `json:"createdAt`
	UserId    int    `json:"userId"`
	EventId   int    `json:"eventId"`
}
