package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response_handler"
	"api/src/security"
	"encoding/json"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		response_handler.ErrorHandler(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	savedUser, err := repository.SearchByEmail(user.Email)
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(savedUser.Password, user.Password); err != nil {
		response_handler.ErrorHandler(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.CreateToken(savedUser.ID)
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	w.Write([]byte(token))
}
