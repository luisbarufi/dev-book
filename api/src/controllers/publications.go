package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response_handler"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePublication(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.ExtractUserId(r)
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusUnauthorized, err)
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publication models.Publication
	if err = json.Unmarshal(requestBody, &publication); err != nil {
		response_handler.ErrorHandler(w, http.StatusBadRequest, err)
		return
	}

	publication.Author_id = userId

	if err = publication.PrepareValidation(); err != nil {
		response_handler.ErrorHandler(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationsRepository(db)
	publication.ID, err = repository.Create(publication)
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	response_handler.JSON(w, http.StatusCreated, publication)
}

func GetPublications(w http.ResponseWriter, r *http.Request) {

}

func ShowPublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	publicationId, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationsRepository(db)
	publication, err := repository.FindById(publicationId)
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	if publication.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response_handler.JSON(w, http.StatusOK, publication)
}

func UpdatePublication(w http.ResponseWriter, r *http.Request) {

}

func DeletePublication(w http.ResponseWriter, r *http.Request) {

}
