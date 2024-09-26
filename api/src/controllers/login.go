package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responseHandler"
	"api/src/security"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responseHandler.ErrorHandler(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	savedUser, err := repository.SearchByEmail(user.Email)
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(savedUser.Password, user.Password); err != nil {
		responseHandler.ErrorHandler(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.CreateToken(savedUser.ID)
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	userId := strconv.FormatUint(savedUser.ID, 10)

	responseHandler.JSON(w, http.StatusOK, models.AuthData{ID: userId, Token: token})
}
