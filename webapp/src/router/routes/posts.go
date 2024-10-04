package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var postsRoutes = []Route{
	{
		URI:                    "/posts",
		Method:                 http.MethodPost,
		Function:               controllers.CreatePost,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{postId}/like",
		Method:                 http.MethodPost,
		Function:               controllers.LikePost,
		RequiresAuthentication: true,
	},
}
