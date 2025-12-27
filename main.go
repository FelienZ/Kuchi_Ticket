package main

import (
	"Go_Ticket/handler"
	"fmt"
	"net/http"
)

var router = http.NewServeMux()

func main() {
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: router,
	}
	router.HandleFunc("/tickets", handler.GetTickets)
	router.HandleFunc("/ticket/{id}", handler.GetTicketById)
	fmt.Printf("Server Run at: http://%s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
