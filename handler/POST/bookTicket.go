package handler

import (
	"Go_Ticket/data"
	"Go_Ticket/model"
	"encoding/json"
	"net/http"
	"strconv"
)

// alias booking
func OrderTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// ticket := &data.Tickets

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid Param",
		})
		return
	}

	var orderData model.Booking

	errOrder := json.NewDecoder(r.Body).Decode(&orderData)
	if errOrder != nil || orderData.UserId == 0 || orderData.EventId == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid Body Request",
		})
		return
	}

	userFound := false
	for _, user := range data.User {
		if user.Id == orderData.UserId {
			userFound = true
			break
		}
	}
	if !userFound {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "User Tidak Ditemukan",
		})
		return
	}

	eventFound := false
	for _, event := range data.Events {
		if event.Id == orderData.EventId {
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
	if orderData.Qty <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Minimal Pemesanan 1 buah",
		})
		return
	}

	for i := range data.Events {
		target := data.Events[i].Id
		if target == id {
			if data.Events[i].Stock <= 0 {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Tiket Habis",
				})
				return
			}
			if orderData.Qty > data.Events[i].Stock {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Stok Tiket Tidak Cukup",
				})
				return
			}
			data.Events[i].Stock = data.Events[i].Stock - orderData.Qty
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]any{
				"message": "Berhasil Order",
				"stok":    data.Events[i].Stock,
			})
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Tiket Tidak Ditemukan",
	})
}
