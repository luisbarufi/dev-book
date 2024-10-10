package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responseHandler"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

func RenderLoginView(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

	utils.ExecuteTemplate(w, "login.html", nil)
}

func RenderUsersRegistrationView(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "registration.html", nil)
}

func RenderHomeView(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/posts", config.ApiUrl)
	response, err := requests.SendAuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		responseHandler.JSON(w, http.StatusInternalServerError, responseHandler.ApiErr{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responseHandler.HandleStatusCodeError(w, response)
		return
	}

	var posts []models.Post
	if err := json.NewDecoder(response.Body).Decode(&posts); err != nil {
		responseHandler.JSON(w, http.StatusUnprocessableEntity, responseHandler.ApiErr{Err: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecuteTemplate(w, "home.html", struct {
		Posts  []models.Post
		UserId uint64
	}{
		Posts:  posts,
		UserId: userId,
	})
}

func RenderEditPostView(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postId, err := strconv.ParseInt(params["postId"], 10, 64)
	if err != nil {
		responseHandler.JSON(w, http.StatusBadRequest, responseHandler.ApiErr{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.ApiUrl, postId)
	response, err := requests.SendAuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		responseHandler.JSON(w, http.StatusInternalServerError, responseHandler.ApiErr{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responseHandler.HandleStatusCodeError(w, response)
		return
	}

	var post models.Post
	if err := json.NewDecoder(response.Body).Decode(&post); err != nil {
		responseHandler.JSON(w, http.StatusUnprocessableEntity, responseHandler.ApiErr{Err: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "update-post.html", post)
}

func RenderSearchUsersView(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	url := fmt.Sprintf("%s/users?user=%s", config.ApiUrl, nameOrNick)

	response, err := requests.SendAuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		responseHandler.JSON(w, http.StatusInternalServerError, responseHandler.ApiErr{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responseHandler.HandleStatusCodeError(w, response)
		return
	}

	var users []models.User
	if err := json.NewDecoder(response.Body).Decode(&users); err != nil {
		responseHandler.JSON(w, http.StatusUnprocessableEntity, responseHandler.ApiErr{Err: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "users.html", users)
}

func LoadUserProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responseHandler.JSON(w, http.StatusBadRequest, responseHandler.ApiErr{Err: err.Error()})
		return
	}

	user, err := models.UserData(userId, r)
	if err != nil {
		responseHandler.JSON(w, http.StatusInternalServerError, responseHandler.ApiErr{Err: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	loggedUserId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecuteTemplate(w, "user.html", struct {
		User         models.User
		LoggedUserId uint64
	}{
		User:         user,
		LoggedUserId: loggedUserId,
	})
}
