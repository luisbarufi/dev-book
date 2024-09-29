package controllers

import (
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/requests"
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

	fmt.Println(response.StatusCode, err)

	utils.ExecuteTemplate(w, "home.html", nil)
}
