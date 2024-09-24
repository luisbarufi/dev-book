package response_handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type ApiErr struct {
	Err string `json:"error"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func HandleStatusCodeError(w http.ResponseWriter, r *http.Response) {
	var err ApiErr
	json.NewDecoder(r.Body).Decode(&err)
	JSON(w, r.StatusCode, err)
}
