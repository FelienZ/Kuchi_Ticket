package handler

import (
	"Go_Ticket/data"
	"Go_Ticket/model"
	"encoding/json"
	"net/http"
)

func GetTickets(w http.ResponseWriter, r *http.Request) {
	tickets := data.Tickets
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	if len(tickets) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"data": "Tiket Sedang Kosong",
		})
		return
	}
	response := map[string][]model.Ticket{
		"data": tickets,
	}
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(response)
}
