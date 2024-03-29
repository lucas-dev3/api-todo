package handlers

import (
	"api-postgres/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	todos, err := models.GetAll()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
