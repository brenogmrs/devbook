package router

import "github.com/gorilla/mux"

// Generate() will return a router with all routes setup
func Generate() *mux.Router {
	return mux.NewRouter()
}
