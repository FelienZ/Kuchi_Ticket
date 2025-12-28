package main

import (
	DEL "Go_Ticket/handler/DELETE"
	GET "Go_Ticket/handler/GET"
	POST "Go_Ticket/handler/POST"
	PUT "Go_Ticket/handler/PUT"
	"fmt"
	"net/http"
)

var router = http.NewServeMux()

func main() {
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: router,
	}

	// Route Event
	router.HandleFunc("/events", GET.GetEvents)
	router.HandleFunc("/event/{id}", GET.GetEventById)
	router.HandleFunc("POST /event", POST.CreateEvent)
	router.HandleFunc("PUT /event/{id}", PUT.UpdateEvent)
	router.HandleFunc("DELETE /event/{id}", DEL.DeleteEvent)
	// Route User
	router.HandleFunc("/users", GET.GetUsers)
	router.HandleFunc("POST /user", POST.CreateUser)
	router.HandleFunc("PUT /user/{id}", PUT.UpdateUser)
	router.HandleFunc("DELETE /user/{id}", DEL.DeleteUser)

	// Route Booking
	router.HandleFunc("POST /event/{id}/booking", POST.OrderTicket)
	router.HandleFunc("PUT /event/{eventId}/booking/{bookingId}", PUT.UpdateBooking)
	router.HandleFunc("DELETE /event/{eventId}/booking/{bookingId}", DEL.DeleteBooking)
	fmt.Printf("Server Run at: http://%s", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
