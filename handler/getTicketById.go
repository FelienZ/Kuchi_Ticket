package handler

import (
	"Go_Ticket/data"
	"Go_Ticket/model"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetTicketById(w http.ResponseWriter, r *http.Request) {
	tickets := data.Tickets
	w.Header().Set("Content-Type", "application/json")

	// id := r.URL.Query().Get("id") //biasanya
	id, err := strconv.Atoi(r.PathValue("id"))
	isExist := false
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Masukkan Id dengan Benar",
		})
		return
	}
	for i := 0; i < len(tickets); i++ {
		val, _ := strconv.Atoi(tickets[i].Id)
		if id == val {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]model.Ticket{
				"data": tickets[i],
			})
			isExist = true
			break
		}
	}
	if !isExist {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Tiket Tidak Ditemukan",
		})
	}
}
