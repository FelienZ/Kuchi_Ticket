package handler

import (
	"Go_Ticket/data"
	"encoding/json"
	"net/http"
	"strconv"
)

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid ID Parameter",
		})
		return
	}

	// validasi, lalu hapus event
	for i, event := range data.Events {
		if event.Id == id {
			data.Events = append(data.Events[:i], data.Events[i+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Berhasil Menghapus Event",
			})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Event Tidak Ditemukan",
	})
}
