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

func Create(w http.ResponseWriter, r *http.Request) {
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

	var post models.Post
	if err = json.Unmarshal(requestBody, &post); err != nil {
		response_handler.ErrorHandler(w, http.StatusBadRequest, err)
		return
	}

	post.Author_id = userId

	if err = post.PrepareValidation(); err != nil {
		response_handler.ErrorHandler(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	post.ID, err = repository.Create(post)
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	response_handler.JSON(w, http.StatusCreated, post)
}

func List(w http.ResponseWriter, r *http.Request) {

}

func Show(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postId, err := strconv.ParseUint(params["postId"], 10, 64)
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

	repository := repositories.NewPostsRepository(db)
	post, err := repository.FindById(postId)
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	if post.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response_handler.JSON(w, http.StatusOK, post)
}

func Update(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
