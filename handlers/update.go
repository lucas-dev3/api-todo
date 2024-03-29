package handlers

import (
	"api-postgres/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error parsing id: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var todo models.Todo
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Update(int64(id), todo)
	if err != nil {
		log.Printf("Error updating todo: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error updating todo: more than one row affected")
	}

	resp := map[string]any{
		"Error":   false,
		"Message": fmt.Sprintf("Todo updated successfully with id %v", id),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
