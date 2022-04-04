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
		RequireAuth: false,
	},
	{
		URI:         "/user/{userId}",
		Method:      http.MethodGet,
		Function:    controllers.GetUserById,
		RequireAuth: false,
	},
	{
		URI:         "/user/{userId}",
		Method:      http.MethodPut,
		Function:    controllers.UpdateUser,
		RequireAuth: false,
	},
	{
		URI:         "/user/{userId}",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteUser,
		RequireAuth: false,
	},
}
