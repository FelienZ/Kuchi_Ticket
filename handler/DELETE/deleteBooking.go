package handler

import (
	"Go_Ticket/data"
	"encoding/json"
	"net/http"
	"strconv"
)

func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	eventId, errEventId := strconv.Atoi(r.PathValue("eventId"))
	_, errBookingId := strconv.Atoi(r.PathValue("bookingId"))

	if errEventId != nil || errBookingId != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid ID Parameter",
		})
		return
	}

	// Validasi event ada
	eventFound := false
	for _, event := range data.Events {
		if event.Id == eventId {
			eventFound = true
			break
		}
	}

	if !eventFound {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Event Tidak Ditemukan",
		})
		return
	}
}
