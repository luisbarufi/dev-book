package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var login = []Route{
	{
		URI:                    "/",
		Method:                 http.MethodGet,
		Function:               controllers.RenderLoginView,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/login",
		Method:                 http.MethodGet,
		Function:               controllers.RenderLoginView,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/login",
		Method:                 http.MethodPost,
		Function:               controllers.Login,
		RequiresAuthentication: false,
	},
}
