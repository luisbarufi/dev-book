package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var users = []Route{
	{
		URI:                    "/create-user",
		Method:                 http.MethodGet,
		Function:               controllers.RenderUsersRegistrationView,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/search-users",
		Method:                 http.MethodGet,
		Function:               controllers.RenderSearchUsersView,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodGet,
		Function:               controllers.LoadUserProfile,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userId}/unfollow",
		Method:                 http.MethodPost,
		Function:               controllers.Unfollow,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userId}/follow",
		Method:                 http.MethodPost,
		Function:               controllers.Follow,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/profile",
		Method:                 http.MethodGet,
		Function:               controllers.RenderLoggedUserProfile,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/edit-user",
		Method:                 http.MethodGet,
		Function:               controllers.RenderUserEditView,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/edit-user",
		Method:                 http.MethodPut,
		Function:               controllers.EditUser,
		RequiresAuthentication: true,
	},
}
