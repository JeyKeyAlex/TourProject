package endpoint

import (
	"encoding/json"
	"github.com/JeyKeyAlex/TourProject/internal/entities"
	"net/http"
	"strconv"
)

func WriteJSON(w http.ResponseWriter, statusCode int, msg any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	var resp any

	switch v := msg.(type) {
	case error:
		resp = map[string]interface{}{"error": v.Error()}
	case string:
		resp = map[string]interface{}{"nextDate": v}
	case int64:
		resp = map[string]interface{}{"id": strconv.FormatInt(v, 10)}
	case []entities.User:
		resp = map[string]interface{}{"users": v}
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
