package routes

import (
	"api/src/controllers"
	"net/http"
)

var postsRoutes = []Route{
	{
		URI:         "/post",
		Method:      http.MethodPost,
		Function:    controllers.CreatePost,
		RequireAuth: true,
	},
	{
		URI:         "/post",
		Method:      http.MethodGet,
		Function:    controllers.GetAllPosts,
		RequireAuth: true,
	},
	{
		URI:         "/post/{postID}",
		Method:      http.MethodGet,
		Function:    controllers.GetPostById,
		RequireAuth: true,
	},
	{
		URI:         "/post/{postID}",
		Method:      http.MethodPut,
		Function:    controllers.UpdatePost,
		RequireAuth: true,
	},
	{
		URI:         "/post/{postID}",
		Method:      http.MethodDelete,
		Function:    controllers.DeletePost,
		RequireAuth: true,
	},
	{
		URI:         "/user/{userID}/posts",
		Method:      http.MethodGet,
		Function:    controllers.GetUserPosts,
		RequireAuth: true,
	},
	{
		URI:         "/post/{postID}/like",
		Method:      http.MethodPost,
		Function:    controllers.LikePost,
		RequireAuth: true,
	},
}
