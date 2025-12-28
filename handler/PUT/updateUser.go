package handler

import (
	"Go_Ticket/data"
	"Go_Ticket/model"
	"encoding/json"
	"net/http"
	"strconv"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid ID Parameter",
		})
		return
	}

	var userData model.User
	errDecode := json.NewDecoder(r.Body).Decode(&userData)
	if errDecode != nil || userData.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid Body Request",
		})
		return
	}

	// Cari user berdasarkan ID
	for i, user := range data.User {
		if user.Id == id {
			data.User[i].Name = userData.Name
			data.User[i].Role = userData.Role
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]any{
				"message": "User berhasil diupdate",
				"data":    data.User[i],
			})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User Tidak Ditemukan",
	})
}
