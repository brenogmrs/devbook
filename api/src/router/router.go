package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate() will return a router with all routes setup
func Generate() *mux.Router {
	r := mux.NewRouter()

	return routes.Config(r)
}
