package handler

import (
	"Go_Ticket/data"
	"Go_Ticket/model"
	"encoding/json"
	"net/http"
	"strconv"
)

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

	var orderData model.Order
	errOrder := json.NewDecoder(r.Body).Decode(&orderData)
	if errOrder != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid Body Request",
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

	for i := range data.Tickets {
		target, _ := strconv.Atoi(data.Tickets[i].Id)
		if target == id {
			if data.Tickets[i].Stock <= 0 {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Tiket Habis",
				})
				return
			}
			if orderData.Qty > data.Tickets[i].Stock {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Stok Tiket Tidak Cukup",
				})
				return
			}
			data.Tickets[i].Stock = data.Tickets[i].Stock - orderData.Qty
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]any{
				"message": "Berhasil Order",
				"stok":    data.Tickets[i].Stock,
			})
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Tiket Tidak Ditemukan",
	})
}
