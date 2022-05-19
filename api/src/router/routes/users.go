package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:         "/user",
		Method:      http.MethodPost,
		Function:    controllers.CreateUser,
		RequireAuth: false,
	},
	{
		URI:         "/user",
		Method:      http.MethodGet,
		Function:    controllers.GetAllUsers,
		RequireAuth: true,
	},
	{
		URI:         "/user/{userID}",
		Method:      http.MethodGet,
		Function:    controllers.GetUserById,
		RequireAuth: true,
	},
	{
		URI:         "/user/{userID}",
		Method:      http.MethodPut,
		Function:    controllers.UpdateUser,
		RequireAuth: true,
	},
	{
		URI:         "/user/{userID}",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteUser,
		RequireAuth: true,
	},
	{
		URI:         "/user/{userID}/follow",
		Method:      http.MethodPost,
		Function:    controllers.FollowUser,
		RequireAuth: true,
	},
}
