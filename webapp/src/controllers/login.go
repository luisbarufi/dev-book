package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"webapp/src/response_handler"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		response_handler.JSON(w, http.StatusBadRequest, response_handler.ApiErr{Err: err.Error()})
		return
	}

	response, err := http.Post("http://localhost:3333/login", "application/json", bytes.NewBuffer(user))
	if err != nil {
		response_handler.JSON(w, http.StatusInternalServerError, response_handler.ApiErr{Err: err.Error()})
		return
	}

	token, _ := io.ReadAll(response.Body)
	fmt.Println(response.StatusCode, string(token))
}
