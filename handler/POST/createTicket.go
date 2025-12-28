package handler

import (
	"Go_Ticket/data"
	"Go_Ticket/model"
	"encoding/json"
	"net/http"
)

// atmin yang bikin tiket nanti
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newData model.Event
	err := json.NewDecoder(r.Body).Decode(&newData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Masukkan Data yang Valid",
		})
		return
	}
	// langsung rujuk ke data dummy agar push jangan copy atau dengan pointer
	newId := len(data.Events) + 1
	newTicket := model.Event{
		Id:    (newId),
		Title: newData.Title,
		Price: newData.Price,
		Stock: newData.Stock,
	}
	if newData.Price <= 0 || newData.Stock < 0 || newData.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Masukkan Data yang Valid",
		})
		return
	}
	data.Events = append(data.Events, newTicket)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"message": "Berhasil Menambahkan",
		"data":    newTicket,
	})
}
