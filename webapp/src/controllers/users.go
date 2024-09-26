package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"webapp/src/responseHandler"
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
		responseHandler.JSON(w, http.StatusBadRequest, responseHandler.ApiErr{Err: err.Error()})
		return
	}

	response, err := http.Post("http://localhost:3333/users", "application/json", bytes.NewBuffer(user))
	if err != nil {
		responseHandler.JSON(w, http.StatusInternalServerError, responseHandler.ApiErr{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responseHandler.HandleStatusCodeError(w, response)
		return
	}

	responseHandler.JSON(w, response.StatusCode, nil)
}
