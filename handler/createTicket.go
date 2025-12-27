package handler

import (
	"Go_Ticket/data"
	"fmt"
	"net/http"
)

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	tickets := data.Tickets
	w.Header().Set("Content-Type", "application/json")

	fmt.Println(tickets)
}
