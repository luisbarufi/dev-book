package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/requests"
	"webapp/src/responseHandler"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	post, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})
	if err != nil {
		responseHandler.JSON(w, http.StatusBadRequest, responseHandler.ApiErr{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts", config.ApiUrl)
	response, err := requests.SendAuthenticatedRequest(r, http.MethodPost, url, bytes.NewBuffer(post))
	if err != nil {
		responseHandler.JSON(w, http.StatusInternalServerError, responseHandler.ApiErr{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responseHandler.HandleStatusCodeError(w, response)
		return
	}

	responseHandler.JSON(w, response.StatusCode, nil)
}
