package handler

import (
	"Go_Ticket/model"
	"encoding/json"
	"net/http"
	"strconv"
)

func UpdateBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, errEventId := strconv.Atoi(r.PathValue("eventId"))
	_, errBookingId := strconv.Atoi(r.PathValue("bookingId"))

	if errEventId != nil || errBookingId != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid ID Parameter",
		})
		return
	}

	var bookingData model.Booking
	errDecode := json.NewDecoder(r.Body).Decode(&bookingData)
	if errDecode != nil || bookingData.UserId == 0 || bookingData.EventId == 0 || bookingData.Qty <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid Body Request",
		})
		return
	}
}
