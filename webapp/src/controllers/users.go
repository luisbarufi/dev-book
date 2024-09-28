package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
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

	url := fmt.Sprintf("%s/users", config.ApiUrl)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(user))
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
