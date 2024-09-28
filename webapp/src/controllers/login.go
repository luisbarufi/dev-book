package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/responseHandler"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		responseHandler.JSON(w, http.StatusBadRequest, responseHandler.ApiErr{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/login", config.ApiUrl)
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

	var authData models.AuthData
	if err = json.NewDecoder(response.Body).Decode(&authData); err != nil {
		responseHandler.JSON(w, http.StatusUnprocessableEntity, responseHandler.ApiErr{Err: err.Error()})
		return
	}

	if err = cookies.Save(w, authData.ID, authData.Token); err != nil {
		responseHandler.JSON(w, http.StatusUnprocessableEntity, responseHandler.ApiErr{Err: err.Error()})
		return
	}

	responseHandler.JSON(w, http.StatusOK, nil)
}
