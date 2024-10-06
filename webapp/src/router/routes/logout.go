package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var logout = Route{
	URI:                    "/logout",
	Method:                 http.MethodGet,
	Function:               controllers.Logout,
	RequiresAuthentication: true,
}
