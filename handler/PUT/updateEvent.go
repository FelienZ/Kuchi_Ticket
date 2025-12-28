package handler

import (
	"Go_Ticket/data"
	"Go_Ticket/model"
	"encoding/json"
	"net/http"
	"strconv"
)

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid ID Parameter",
		})
		return
	}

	var eventData model.Event
	errDecode := json.NewDecoder(r.Body).Decode(&eventData)
	if errDecode != nil || eventData.Title == "" || eventData.Price <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid Body Request",
		})
		return
	}

	// Cari event berdasarkan id, update
	for i, event := range data.Events {
		if event.Id == id {
			data.Events[i].Title = eventData.Title
			data.Events[i].Price = eventData.Price
			data.Events[i].Stock = eventData.Stock
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]any{
				"message": "Event berhasil diupdate",
				"data":    data.Events[i],
			})
			return
		}
	}

	// !ditemukan
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Event Tidak Ditemukan",
	})
}
