package routes

import (
	"api/src/controllers"
	"net/http"
)

var postsRoutes = []Route{
	{
		URI:                    "/posts",
		Method:                 http.MethodPost,
		Function:               controllers.Create,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts",
		Method:                 http.MethodGet,
		Function:               controllers.List,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{postId}",
		Method:                 http.MethodGet,
		Function:               controllers.Show,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{postId}",
		Method:                 http.MethodPut,
		Function:               controllers.Update,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{postId}",
		Method:                 http.MethodDelete,
		Function:               controllers.Delete,
		RequiresAuthentication: true,
	},
}
