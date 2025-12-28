package handler

import (
	"Go_Ticket/data"
	"Go_Ticket/model"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newData model.User
	err := json.NewDecoder(r.Body).Decode(&newData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Masukkan Data yang Valid",
		})
		return
	}
	newId := len(data.User) + 1
	newUser := model.User{
		Id:   newId,
		Name: newData.Name,
		Role: newData.Role,
	}
	if newUser.Name == "" || newUser.Role == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Masukkan Data yang Valid",
		})
		return
	}
	data.User = append(data.User, newUser)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"message": "Berhasil Menambahkan",
		"data":    newUser,
	})
}
