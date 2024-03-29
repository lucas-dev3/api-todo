package handlers

import (
	"api-postgres/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, http.StatusText((http.StatusInternalServerError)), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)

	var response map[string]any

	if err != nil {
		response = map[string]any{"Status": true, "Message": fmt.Sprintf("Error inserting todo: %v", err)}
	} else {
		response = map[string]any{"Status": true, "Message": fmt.Sprintf("Todo inserted successfully with id %v", id)}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
