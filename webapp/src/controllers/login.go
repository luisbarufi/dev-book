package controllers

import "net/http"

func RenderLoginView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Tela de Login"))
}
