package routes

import (
	"net/http"
	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

func Configure(router *mux.Router) *mux.Router {
	routes := login
	routes = append(routes, usersRoutes...)
	routes = append(routes, home)
	routes = append(routes, postsRoutes...)

	for _, route := range routes {
		if route.RequiresAuthentication {
			router.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Function)),
			).Methods(route.Method)

		} else {
			router.HandleFunc(route.URI,
				middlewares.Logger(route.Function),
			).Methods(route.Method)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
