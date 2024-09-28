package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var home = Route{
	URI:                    "/home",
	Method:                 http.MethodGet,
	Function:               controllers.RenderHomeView,
	RequiresAuthentication: true,
}
