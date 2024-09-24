package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"webapp/src/response_handler"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"nick":     r.FormValue("nick"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		response_handler.JSON(w, http.StatusBadRequest, response_handler.ApiErr{Err: err.Error()})
		return
	}

	response, err := http.Post("http://localhost:3333/users", "application/json", bytes.NewBuffer(user))
	if err != nil {
		response_handler.JSON(w, http.StatusInternalServerError, response_handler.ApiErr{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		response_handler.HandleStatusCodeError(w, response)
		return
	}

	response_handler.JSON(w, response.StatusCode, nil)
}
