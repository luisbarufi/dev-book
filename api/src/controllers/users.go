package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response_handler"
	"api/src/security"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
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

	if err = user.PrepareValidation("register"); err != nil {
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
	user.ID, err = repository.Create(user)
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	response_handler.JSON(w, http.StatusCreated, user)
}

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	searchParameter := strings.ToLower(r.URL.Query().Get("user"))
	db, err := database.Connect()
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	users, err := repository.Search(searchParameter)
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	response_handler.JSON(w, http.StatusOK, users)
}

func ShowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
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

	repository := repositories.NewUserRepository(db)
	user, err := repository.FindById(userId)
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response_handler.JSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusBadRequest, err)
		return
	}

	userIdInToken, err := auth.ExtractUserId(r)
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusUnauthorized, err)
		return
	}

	if userId != userIdInToken {
		response_handler.ErrorHandler(w, http.StatusForbidden, errors.New("não é possível atualizar um usuário que não seja o seu"))
		return
	}

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

	if err = user.PrepareValidation("edit"); err != nil {
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
	if err = repository.Update(userId, user); err != nil {
		response_handler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	response_handler.JSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusBadRequest, err)
		return
	}

	userIdInToken, err := auth.ExtractUserId(r)
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusUnauthorized, err)
		return
	}

	if userId != userIdInToken {
		response_handler.ErrorHandler(w, http.StatusForbidden, errors.New("não é possível deletar um usuário que não seja o seu"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if err = repository.Delete(userId); err != nil {
		response_handler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	response_handler.JSON(w, http.StatusNoContent, nil)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, err := auth.ExtractUserId(r)
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusBadRequest, err)
		return
	}

	if followerId == userId {
		response_handler.ErrorHandler(w, http.StatusForbidden, errors.New("não é possível seguir você mesmo"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if err := repository.Follow(userId, followerId); err != nil {
		response_handler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	response_handler.JSON(w, http.StatusNoContent, nil)
}

func UnFollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, err := auth.ExtractUserId(r)
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusBadRequest, err)
		return
	}

	if followerId == userId {
		response_handler.ErrorHandler(w, http.StatusForbidden, errors.New("não é possível parar de seguir você mesmo"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if err := repository.UnFollow(userId, followerId); err != nil {
		response_handler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	response_handler.JSON(w, http.StatusNoContent, nil)
}

func SearchFollowers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
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

	repository := repositories.NewUserRepository(db)
	followers, err := repository.SearchFollowers(userId)
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	response_handler.JSON(w, http.StatusOK, followers)
}

func SearchFollowing(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
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

	repository := repositories.NewUserRepository(db)
	users, err := repository.SearchFollowing(userId)
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	response_handler.JSON(w, http.StatusOK, users)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userIdInToken, err := auth.ExtractUserId(r)
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusBadRequest, err)
		return
	}

	if userIdInToken != userId {
		response_handler.ErrorHandler(w, http.StatusForbidden, errors.New("não é possível atualizar a senha de um usuário que não é o seu"))
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusUnprocessableEntity, err)
		return
	}

	var password models.Password

	if err = json.Unmarshal(requestBody, &password); err != nil {
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
	passwordSaved, err := repository.GetSavedPassword(userId)
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(passwordSaved, password.CurrentPassword); err != nil {
		response_handler.ErrorHandler(w, http.StatusUnauthorized, errors.New("a senha atual não condiz coma a que está salva no banco"))
		return
	}

	hashedPassword, err := security.Hash(password.NewPassword)
	if err != nil {
		response_handler.ErrorHandler(w, http.StatusBadRequest, err)
		return
	}

	if err = repository.UpdatePassword(userId, string(hashedPassword)); err != nil {
		response_handler.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	response_handler.JSON(w, http.StatusNoContent, nil)
}
