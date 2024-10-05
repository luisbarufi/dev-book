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
	{
		URI:                    "/posts/{postId}/dislike",
		Method:                 http.MethodPost,
		Function:               controllers.DisLikePost,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{postId}/edit",
		Method:                 http.MethodGet,
		Function:               controllers.RenderEditPostView,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{postId}",
		Method:                 http.MethodPut,
		Function:               controllers.EditPost,
		RequiresAuthentication: true,
	},
}
