package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responseHandler"
	"webapp/src/utils"
)

func RenderLoginView(w http.ResponseWriter, r *http.Request) {
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

	utils.ExecuteTemplate(w, "home.html", posts)
}
