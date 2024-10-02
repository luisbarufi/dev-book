package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responseHandler"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.ExtractUserId(r)
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusUnauthorized, err)
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(requestBody, &post); err != nil {
		responseHandler.ErrorHandler(w, http.StatusBadRequest, err)
		return
	}

	post.AuthorId = userId

	if err = post.PrepareValidation(); err != nil {
		responseHandler.ErrorHandler(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	post.ID, err = repository.Create(post)
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	responseHandler.JSON(w, http.StatusCreated, post)
}

func ListPosts(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.ExtractUserId(r)
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	posts, err := repository.ListPosts(userId)
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	responseHandler.JSON(w, http.StatusCreated, posts)
}

func ShowPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postId, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	post, err := repository.FindById(postId)
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	if post.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	responseHandler.JSON(w, http.StatusOK, post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.ExtractUserId(r)
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)

	postId, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	postSaved, err := repository.FindById(postId)
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	if postSaved.AuthorId != userId {
		responseHandler.ErrorHandler(w, http.StatusForbidden, errors.New("não é possível atualizar uma publicação que não seja sua"))
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(requestBody, &post); err != nil {
		responseHandler.ErrorHandler(w, http.StatusBadRequest, err)
		return
	}

	if err = post.PrepareValidation(); err != nil {
		responseHandler.ErrorHandler(w, http.StatusBadRequest, err)
		return
	}

	if err = repository.UpdatePost(postId, post); err != nil {
		responseHandler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	responseHandler.JSON(w, http.StatusNoContent, nil)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.ExtractUserId(r)
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)

	postId, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	postSaved, err := repository.FindById(postId)
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	if postSaved.AuthorId != userId {
		responseHandler.ErrorHandler(w, http.StatusForbidden, errors.New("não é possível deletar uma publicação que não seja sua"))
		return
	}

	if err = repository.DeletePost(postId); err != nil {
		responseHandler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	responseHandler.JSON(w, http.StatusNoContent, nil)
}

func ListPostsByUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	posts, err := repository.ListPostsByUser(userId)
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	responseHandler.JSON(w, http.StatusOK, posts)
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postId, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	if err = repository.LikePost(postId); err != nil {
		responseHandler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	responseHandler.JSON(w, http.StatusNoContent, nil)
}

func DisLikePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postId, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responseHandler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	if err = repository.DisLikePost(postId); err != nil {
		responseHandler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	responseHandler.JSON(w, http.StatusNoContent, nil)
}
